package oauth

import (
	"context"
	"fmt"
	"oauth2/internal/core"
	"oauth2/internal/model"
	proto "oauth2/rpc/oauth2"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

const (
	_defaltAppKey = "normal"
)

// Oauth 授权
type Oauth struct {
}

// Authorize 应用授权
func (h *Oauth) Authorize(ctx context.Context, req *proto.AuthorizeReq) (*proto.AuthorizeResp, error) {
	if len(req.AppId) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	oauthinfo, err := getOneOauth2Tbl(req.AppId, req.AppKey)
	if err != nil {
		return nil, err
	}
	//--------------授权校验client
	if oauthinfo.StrictSign { // 需要验签
		if !verifyToken(req.Token, fmt.Sprintf("%v%v%v%v", req.AppId, req.AppKey, req.Timestamp, oauthinfo.AppSecret)) { // 验签失败
			return nil, message.GetError(message.TokenCheckError)
		}
	}
	// -----------------------end
	if len(req.AppId) == 0 {
		req.AppId = oauthinfo.AppID
	}

	userName := req.UserName
	if len(userName) == 0 {
		userName = req.AppId
	}
	// 校验成功，生成 reflash/access_token
	token, err := newToken(oauthinfo, userName, req.TokenType)
	if err != nil {
		return nil, err
	}

	return &proto.AuthorizeResp{
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		AccessExpireTime:  token.AccessExpireTime,
		RefreshExpireTime: token.RefreshExpireTime,
	}, nil
}

// CheckToken 校验token，并获取详细信息
func (h *Oauth) CheckToken(ctx context.Context, req *proto.CheckTokenReq) (*proto.CheckTokenResp, error) {
	if len(req.Token) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	// 缓存中获取
	tmp := GetCacheToken(req.Token, _prefixAccess)
	if tmp == nil { // 未找到
		tmp = &TokenCache{}
	}

	if len(tmp.Token) == 0 { // 数据库找
		orm := core.Dao.GetDBr()
		info, err := model.AccessTokenTblMgr(orm.DB).GetFromAccessToken(req.Token)
		if err != nil {
			if orm.IsNotFound(err) { // 未找到
				return nil, message.GetError(message.NotFindError)
			}

			mylog.Error(err)
			return nil, err
		}
		if info.ID <= 0 { // 未找到
			return nil, message.GetError(message.NotFindError)
		}
		tmp.Token = info.AccessToken
		tmp.ExpireTime = info.Expires.Unix()
		tmp.UserName = info.Username
		tmp.AppID = info.AppID
		tmp.TokenType = info.TokenType
	}

	// 判断过期
	if tmp.ExpireTime <= time.Now().Unix() {
		return nil, message.GetError(message.Overdue)
	}

	// 未过期，直接返回
	return &proto.CheckTokenResp{
		AccessToken:      tmp.Token,
		UserName:         tmp.UserName,
		AccessExpireTime: tmp.ExpireTime,
		AppId:            tmp.AppID,
	}, nil
}

// RefreshToken 刷新token
func (h *Oauth) RefreshToken(ctx context.Context, req *proto.RefreshTokenReq) (*proto.RefreshTokenResp, error) {
	if len(req.Token) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	// 缓存中获取
	tmp := GetCacheToken(req.Token, _prefixRefresh)
	if tmp == nil { // 未找到
		tmp = &TokenCache{}
	}

	if len(tmp.Token) == 0 { // 数据库找
		orm := core.Dao.GetDBr()
		info, err := model.RefreshTokenTblMgr(orm.DB).GetFromRefreshToken(req.Token)
		if err != nil {
			if orm.IsNotFound(err) { // 未找到
				return nil, message.GetError(message.NotFindError)
			}

			mylog.Error(err)
			return nil, err
		}
		if info.ID <= 0 { // 未找到
			return nil, message.GetError(message.NotFindError)
		}
		tmp.Token = info.RefreshToken
		tmp.ExpireTime = info.Expires.Unix()
		tmp.UserName = info.Username
		tmp.AppID = info.AppID
		tmp.TokenType = info.TokenType
	}

	// 判断过期
	if tmp.ExpireTime <= time.Now().Unix() {
		return nil, message.GetError(message.Overdue)
	}

	// 未过期，刷新token

	oauthinfo, err := getOneOauth2Tbl(tmp.AppID, "")
	if err != nil {
		return nil, err
	}

	token, err := newToken(oauthinfo, oauthinfo.Username, tmp.TokenType)
	if err != nil {
		return nil, err
	}
	// todo:删除refreshToken

	return &proto.RefreshTokenResp{
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		AccessExpireTime:  token.AccessExpireTime,
		RefreshExpireTime: token.RefreshExpireTime,
	}, nil
}

// CreateOauth 创建oauth
func (h *Oauth) CreateOauth(ctx context.Context, req *proto.CreateOauthReq) (*proto.CreateOauthResp, error) {
	if len(req.Username) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBw()
	var times [5][0]int
	for range times {
		oauthInfo := model.Oauth2Tbl{
			AppID:           tools.GetRangeNumString(16), // 应用的唯一标识
			AppKey:          tools.GetRandomString(24),   // 公匙
			AppSecret:       tools.GetRandomString(32),   // 私匙
			Username:        req.Username,                // 用户账号
			ExpireTime:      time.Unix(2524579200, 0),    // appid超时时间
			TokenExpireTime: 1000000,                     // token过期时间
			OauthInfo:       req.OauthInfo,
			CreatedBy:       req.Username,
		}
		oauthInfo.Model.CreatedAt = time.Now()

		var count int64
		if err := model.Oauth2TblMgr(orm.Where("app_id = ?", oauthInfo.AppID)).Count(&count).Error; err != nil {
			return nil, err
		}

		if count == 0 {
			if err := model.Oauth2TblMgr(orm.DB).Save(&oauthInfo).Error; err != nil {
				return nil, err
			} else { // 创建成功
				return &proto.CreateOauthResp{
					AppId:     oauthInfo.AppID,  // 应用id
					AppKey:    oauthInfo.AppKey, // 公匙
					AppSecret: oauthInfo.AppSecret,
				}, nil
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil, message.GetError(message.UnknownError)
}

// GetAppList 获取应用信息
func (h *Oauth) GetAppList(ctx context.Context, req *proto.GetAppListReq) (*proto.GetAppListResp, error) {
	if len(req.Username) == 0 && len(req.AppIds) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	resp := &proto.GetAppListResp{}
	orm := core.Dao.GetDBr()

	var list []*model.Oauth2Tbl
	var err error

	if len(req.AppIds) > 0 {
		// offse
		db := orm.Offset(int(req.PageNo * req.PageSize)).Limit(int(req.PageSize)).Order("created_at desc")
		list, err = model.Oauth2TblMgr(db).GetBatchFromAppID(req.AppIds)
		if err != nil {
			return nil, err
		}
	} else {
		// 获取数量
		model.Oauth2TblMgr(orm.Where("username = ?", req.Username)).Count(&(resp.Total))

		// offse
		db := orm.Offset(int(req.PageNo * req.PageSize)).Limit(int(req.PageSize)).Order("created_at desc")
		list, err = model.Oauth2TblMgr(db).GetFromUsername(req.Username)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range list {
		resp.List = append(resp.List, &proto.AppInfo{
			AppId:     v.AppID, // appid
			OauthInfo: v.OauthInfo,
		})
	}
	return resp, nil
}

// DeleteAppList 删除应用
func (h *Oauth) DeleteApp(ctx context.Context, req *proto.DeleteAppReq) (*proto.DeleteAppResp, error) {
	if len(req.Username) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBw()
	num := orm.Where("app_id = ?", req.AppId).Delete(&model.Oauth2Tbl{}).RowsAffected
	return &proto.DeleteAppResp{
		RowsAffected: num,
	}, nil
}
