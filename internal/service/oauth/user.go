package oauth

import (
	"context"
	"fmt"
	"oauth2/internal/core"
	"oauth2/internal/model"
	"strings"
	"time"

	proto "oauth2/rpc/oauth2"
	"rpc/common"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
)

// Login 登录
func (h *Oauth) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp, error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	//--------------验签
	if !verifyToken(req.Token, fmt.Sprintf("%v%v%v", req.Username, req.Password, req.Timestamp)) { // 验签失败
		return nil, message.GetError(message.TokenCheckError)
	}
	// -----------------------end

	orm := core.Dao.GetDBr()
	info, err := model.UserAccountTblMgr(orm.DB).GetFromUsername(req.Username)
	if err != nil {
		if orm.IsNotFound(err) { // 未找到
			return nil, message.GetError(message.UserNameDoNotExist)
		}

		mylog.Error(err)
		return nil, err
	}

	if info.ID <= 0 || !strings.EqualFold(req.Password, info.Password) {
		return nil, message.GetError(message.UserNameDoNotExist)
	}

	if !info.Vaild { // 已无效
		return nil, message.GetError(message.InValidOp)
	}

	oauthinfo := &info.Oauth2Tbl
	if oauthinfo.ID <= 0 { // 使用默认appkey
		oauthinfo, err = getOneOauth2Tbl("", _defaltAppKey)
		if err != nil {
			return nil, err
		}
	}

	// 校验成功，生成 reflash/access_token
	token, err := newToken(oauthinfo, req.Username, "login")
	if err != nil {
		return nil, err
	}

	return &proto.LoginResp{
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		AccessExpireTime:  token.AccessExpireTime,
		RefreshExpireTime: token.RefreshExpireTime,
	}, nil
}

// GetLogInfo 获取用户信息
func (h *Oauth) GetLoginInfo(ctx context.Context, req *proto.GetLoginInfoReq) (*proto.GetLoginInfoResp, error) {
	if len(req.AccessToken) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	resp, err := h.CheckToken(ctx, &proto.CheckTokenReq{
		Token: req.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.UserName) == 0 {
		return nil, message.GetError(message.TokenFailure)
	}

	orm := core.Dao.GetDBr()
	re, err := model.UserInfoTblMgr(orm.DB).GetFromUsername(resp.UserName)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, message.GetError(message.NotFindError)
		}
		mylog.Error(err)
		return nil, err
	}
	if re.ID <= 0 {
		return nil, message.GetError(message.NotFindError)
	}

	return &proto.GetLoginInfoResp{
		Username:         resp.UserName,
		UserInfo:         re.UserInfo,
		AccessExpireTime: resp.AccessExpireTime,
	}, nil
}

// CreateUser 创建用户
func (h *Oauth) CreateUser(ctx context.Context, req *proto.CreateUserReq) (resp *common.Empty, _err error) {
	if len(req.RootName) == 0 || len(req.UserName) == 0 || len(req.Password) == 0 || len(req.UserInfo) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBw()
	tx := orm.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()

	mgr := model.UserAccountTblMgr(tx)
	re, _ := mgr.GetFromUsername(req.UserName)
	if re.ID > 0 {
		return nil, message.GetError(message.UserExisted)
	}

	userInfoTbl := model.UserInfoTbl{
		Username:  req.UserName,
		UserInfo:  req.UserInfo,
		CreatedBy: req.RootName, // 创建者
		CreatedAt: time.Now(),   // 创建时间
	}

	model.UserInfoTblMgr(tx).Save(&userInfoTbl)

	info := model.UserAccountTbl{
		Username:    req.UserName,                  // 用户账号
		Password:    strings.ToUpper(req.Password), // 用户密码
		AppID:       "hainlp",                      // oauth2_tbl表的id(验签id)
		UserInfoID:  userInfoTbl.ID,                // 用户附加信息id
		UserInfoTbl: userInfoTbl,
		RegTime:     time.Now(),   // 注册时间
		RegIP:       req.RegIP,    // 注册ip
		Vaild:       true,         // 是否有效
		CreatedBy:   req.RootName, // 创建者
		CreatedAt:   time.Now(),   // 创建时间
	}
	_err = mgr.Save(&info).Error

	return &common.Empty{}, _err
}

// GetUsers 获取用户列表
func (h *Oauth) GetUsers(ctx context.Context, req *proto.GetUsersReq) (*proto.GetUsersResp, error) {
	if len(req.RootName) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBr()
	mgr := model.UserAccountTblMgr(orm.DB)

	var options []model.Option
	if len(req.UserName) > 0 {
		options = append(options, mgr.WithUsername(req.UserName))
	} else {
		options = append(options, mgr.WithCreatedBy(req.RootName))
	}
	res, err := mgr.GetByOptions(options...)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, message.GetError(message.NotFindError)
		}
		mylog.Error(err)
		return nil, err
	}

	resp := &proto.GetUsersResp{}
	for _, v := range res {
		resp.List = append(resp.List, &proto.UserInfo{
			Username: v.Username, // 用户名
			UserInfo: v.UserInfoTbl.UserInfo,
		})
	}
	return resp, nil
}

// UpdateUser 更新/删除用户
func (h *Oauth) UpdateUser(ctx context.Context, req *proto.UpdateUserReq) (*common.Empty, error) {
	if len(req.RootName) == 0 || len(req.UserName) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBw()
	mgr := model.UserAccountTblMgr(orm.DB)
	account, err := mgr.GetFromUsername(req.UserName)
	if account.ID == 0 || err != nil {
		if account.ID == 0 || orm.IsNotFound(err) {
			return nil, message.GetError(message.NotFindError)
		}
		mylog.Error(err)
		return nil, err
	}

	switch req.Op {
	case 1: // 更新
		if len(req.Password) > 0 { // 更新密码
			err = mgr.Where("username = ?", req.UserName).Updates(map[string]interface{}{"password": strings.ToUpper(req.Password), "updated_by": req.RootName, "updated_at": time.Now()}).Error
			if err != nil {
				return nil, err
			}
		}
		if len(req.UserInfo) > 0 { // 更新用户信息
			err = model.UserInfoTblMgr(orm.Where("username = ?", req.UserName)).Updates(map[string]interface{}{"user_info": req.UserInfo,
				"updated_by": req.RootName, "updated_at": time.Now()}).Error
			if err != nil {
				return nil, err
			}
		}
	case -1: // 删除
		err = mgr.Delete(&account).Error
		if err != nil {
			return nil, err
		}
		err = model.UserInfoTblMgr(orm.DB).Delete(&account.UserInfoTbl).Error
		if err != nil {
			return nil, err
		}
	default:
		return nil, message.GetError(message.ParameterInvalid)
	}

	return &common.Empty{}, nil
}
