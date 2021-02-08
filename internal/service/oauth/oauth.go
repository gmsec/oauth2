package oauth

import (
	"context"
	"fmt"
	"oauth2/internal/core"
	"oauth2/internal/model"
	proto "oauth2/rpc/oauth2"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
)

const (
	_defaltAppKey = "normal"
)

// Oauth 授权
type Oauth struct {
}

// Authorize 应用授权
func (h *Oauth) Authorize(ctx context.Context, req *proto.AuthorizeReq) (*proto.AuthorizeResp, error) {
	if len(req.AppKey) == 0 {
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

	// 校验成功，生成 reflash/access_token
	token, err := newToken(oauthinfo, req.AppId, req.TokenType)
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
		tmp.UserInfo = info.Userinfo
		tmp.AppKey = info.AppKey
		tmp.TokenType = info.TokenType
	}

	// 判断过期
	if tmp.ExpireTime <= time.Now().Unix() {
		return nil, message.GetError(message.Overdue)
	}

	// 未过期，直接返回
	return &proto.CheckTokenResp{
		AccessToken:      tmp.Token,
		UserInfo:         tmp.UserInfo,
		AccessExpireTime: tmp.ExpireTime,
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
		tmp.UserInfo = info.Userinfo
		tmp.AppKey = info.AppKey
		tmp.TokenType = info.TokenType
	}

	// 判断过期
	if tmp.ExpireTime <= time.Now().Unix() {
		return nil, message.GetError(message.Overdue)
	}

	// 未过期，刷新token

	oauthinfo, err := getOneOauth2Tbl("", tmp.AppKey)
	if err != nil {
		return nil, err
	}

	token, err := newToken(oauthinfo, oauthinfo.AppID, tmp.TokenType)
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
