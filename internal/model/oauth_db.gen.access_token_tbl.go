package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AccessTokenTblMgr struct {
	*_BaseMgr
}

// AccessTokenTblMgr open func
func AccessTokenTblMgr(db *gorm.DB) *_AccessTokenTblMgr {
	if db == nil {
		panic(fmt.Errorf("AccessTokenTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AccessTokenTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("access_token_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AccessTokenTblMgr) GetTableName() string {
	return "access_token_tbl"
}

// Get 获取
func (obj *_AccessTokenTblMgr) Get() (result AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AccessTokenTblMgr) Gets() (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AccessTokenTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAccessToken access_token获取 访问令牌
func (obj *_AccessTokenTblMgr) WithAccessToken(accessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = accessToken })
}

// WithTokenType token_type获取 令牌类型
func (obj *_AccessTokenTblMgr) WithTokenType(tokenType string) Option {
	return optionFunc(func(o *options) { o.query["token_type"] = tokenType })
}

// WithAppKey app_key获取 key
func (obj *_AccessTokenTblMgr) WithAppKey(appKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = appKey })
}

// WithUserinfo userinfo获取 用户名
func (obj *_AccessTokenTblMgr) WithUserinfo(userinfo string) Option {
	return optionFunc(func(o *options) { o.query["userinfo"] = userinfo })
}

// WithExpires expires获取 过期时间
func (obj *_AccessTokenTblMgr) WithExpires(expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = expires })
}

// GetByOption 功能选项模式获取
func (obj *_AccessTokenTblMgr) GetByOption(opts ...Option) (result AccessTokenTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AccessTokenTblMgr) GetByOptions(opts ...Option) (results []*AccessTokenTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AccessTokenTblMgr) GetFromID(id int) (result AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_AccessTokenTblMgr) GetBatchFromID(ids []int) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容 访问令牌
func (obj *_AccessTokenTblMgr) GetFromAccessToken(accessToken string) (result AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("access_token = ?", accessToken).Find(&result).Error

	return
}

// GetBatchFromAccessToken 批量唯一主键查找 访问令牌
func (obj *_AccessTokenTblMgr) GetBatchFromAccessToken(accessTokens []string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("access_token IN (?)", accessTokens).Find(&results).Error

	return
}

// GetFromTokenType 通过token_type获取内容 令牌类型
func (obj *_AccessTokenTblMgr) GetFromTokenType(tokenType string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_type = ?", tokenType).Find(&results).Error

	return
}

// GetBatchFromTokenType 批量唯一主键查找 令牌类型
func (obj *_AccessTokenTblMgr) GetBatchFromTokenType(tokenTypes []string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_type IN (?)", tokenTypes).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容 key
func (obj *_AccessTokenTblMgr) GetFromAppKey(appKey string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key = ?", appKey).Find(&results).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找 key
func (obj *_AccessTokenTblMgr) GetBatchFromAppKey(appKeys []string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key IN (?)", appKeys).Find(&results).Error

	return
}

// GetFromUserinfo 通过userinfo获取内容 用户名
func (obj *_AccessTokenTblMgr) GetFromUserinfo(userinfo string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("userinfo = ?", userinfo).Find(&results).Error

	return
}

// GetBatchFromUserinfo 批量唯一主键查找 用户名
func (obj *_AccessTokenTblMgr) GetBatchFromUserinfo(userinfos []string) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("userinfo IN (?)", userinfos).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容 过期时间
func (obj *_AccessTokenTblMgr) GetFromExpires(expires time.Time) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expires = ?", expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量唯一主键查找 过期时间
func (obj *_AccessTokenTblMgr) GetBatchFromExpires(expiress []time.Time) (results []*AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expires IN (?)", expiress).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AccessTokenTblMgr) FetchByPrimaryKey(id int) (result AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByAccessToken primay or index 获取唯一内容
func (obj *_AccessTokenTblMgr) FetchUniqueByAccessToken(accessToken string) (result AccessTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("access_token = ?", accessToken).Find(&result).Error

	return
}
