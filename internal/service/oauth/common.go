package oauth

import (
	"fmt"
	"oauth2/internal/core"
	"oauth2/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// verifyToken 验签token
func verifyToken(token, single string) bool {
	code := tools.Md5Encoder(single)
	return strings.EqualFold(code, token)
}

// GetOne 获取一个key可能有缓存
func getOneOauth2Tbl(appID, appKey string) (*model.Oauth2Tbl, error) {
	out := GetCacheOauth2TblByKey(appID) // 缓存中获取
	if out != nil {
		return out, nil
	}

	orm := core.Dao.GetDBr()
	mgr := model.Oauth2TblMgr(orm.Where("expire_time > ?", time.Now()))

	var options []model.Option
	if len(appID) > 0 {
		options = append(options, mgr.WithAppID(appID))
	}
	if len(appKey) > 0 {
		options = append(options, mgr.WithAppKey(appKey))
	}
	oauthinfo, err := mgr.GetByOption(options...)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, message.GetError(message.NotFindError)
		}
		mylog.Error(err)
		return nil, message.GetError(message.ServerMaintenance)
	}
	if oauthinfo.ID <= 0 {
		return nil, fmt.Errorf(message.NotFindError.String())
	}

	err = AddCacheOauth2TblByKey(appID, &oauthinfo) // 缓存中获取
	if err != nil {
		mylog.Error(err)
	}

	return &oauthinfo, nil
}

func newToken(oauth2Info *model.Oauth2Tbl, username, tokenType string) (*TokenInfo, error) {
	rtoken, err := newRefreshToken(oauth2Info, username, tokenType)
	if err != nil {
		return nil, err
	}
	atoken, err := newAccessToken(oauth2Info, username, tokenType)
	if err != nil {
		return nil, err
	}
	info := &TokenInfo{}
	info.AccessToken = atoken.Token
	info.AccessExpireTime = atoken.ExpireTime
	info.UserName = atoken.UserName
	info.RefreshToken = rtoken.Token
	info.RefreshExpireTime = rtoken.ExpireTime

	return info, nil
}

func newAccessToken(oauth2Info *model.Oauth2Tbl, username, tokenType string) (*TokenCache, error) {
	var err error
	info := model.AccessTokenTbl{
		TokenType: tokenType,        // 令牌类型
		AppID:     oauth2Info.AppID, // key
		Username:  username,         // 用户名
		// Expires :  time.Now().Add(time.Duration(token_expire_time) * time.Second)                    // 过期时间
	}

	orm := core.Dao.GetDBw()
	mgr := model.AccessTokenTblMgr(orm.DB)
	for i := 0; i < 3; i++ {
		info.AccessToken = tools.GetRandomString(32)
		info.Expires = time.Now().Add(time.Duration(oauth2Info.TokenExpireTime) * time.Second)

		_, err = mgr.GetFromAccessToken(info.AccessToken)
		if orm.IsNotFound(err) { // 没找到
			err = mgr.Save(&info).Error
			if err != nil {
				mylog.Error(err)
			} else { // 成功
				break
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	cache := &TokenCache{
		Token:      info.AccessToken,
		ExpireTime: info.Expires.Unix(),
		UserName:   username,
		AppID:      oauth2Info.AppID,
		TokenType:  tokenType,
	}

	// 保存缓存
	err = AddCacheToken(*cache, _prefixAccess)
	return cache, err
}

func newRefreshToken(oauth2Info *model.Oauth2Tbl, username, tokenType string) (*TokenCache, error) {
	var err error
	info := model.RefreshTokenTbl{
		TokenType: tokenType,        // 令牌类型
		AppID:     oauth2Info.AppID, // key
		Username:  username,         // 用户名
		// Expires :  time.Now().Add(time.Duration(token_expire_time) * time.Second)                    // 过期时间
	}

	orm := core.Dao.GetDBw()
	mgr := model.RefreshTokenTblMgr(orm.DB)
	for i := 0; i < 3; i++ {
		info.RefreshToken = tools.GetRandomString(32)
		info.Expires = time.Now().Add(time.Duration(oauth2Info.TokenExpireTime*2) * time.Second)

		_, err = mgr.GetFromRefreshToken(info.RefreshToken)
		if orm.IsNotFound(err) { // 没找到
			err = mgr.Save(&info).Error
			if err != nil {
				mylog.Error(err)
			} else { // 成功
				break
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	cache := &TokenCache{
		Token:      info.RefreshToken,
		ExpireTime: info.Expires.Unix(),
		UserName:   username,
		AppID:      oauth2Info.AppID,
		TokenType:  tokenType,
	}

	// 保存缓存
	err = AddCacheToken(*cache, _prefixRefresh)
	return cache, err
}
