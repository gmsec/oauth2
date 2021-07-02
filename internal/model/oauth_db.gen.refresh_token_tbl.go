package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _RefreshTokenTblMgr struct {
	*_BaseMgr
}

// RefreshTokenTblMgr open func
func RefreshTokenTblMgr(db *gorm.DB) *_RefreshTokenTblMgr {
	if db == nil {
		panic(fmt.Errorf("RefreshTokenTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RefreshTokenTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("refresh_token_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RefreshTokenTblMgr) GetTableName() string {
	return "refresh_token_tbl"
}

// Get 获取
func (obj *_RefreshTokenTblMgr) Get() (result RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RefreshTokenTblMgr) Gets() (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RefreshTokenTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRefreshToken refresh_token获取 刷新令牌
func (obj *_RefreshTokenTblMgr) WithRefreshToken(refreshToken string) Option {
	return optionFunc(func(o *options) { o.query["refresh_token"] = refreshToken })
}

// WithTokenType token_type获取 令牌类型
func (obj *_RefreshTokenTblMgr) WithTokenType(tokenType string) Option {
	return optionFunc(func(o *options) { o.query["token_type"] = tokenType })
}

// WithAppID app_id获取 应用的唯一标识
func (obj *_RefreshTokenTblMgr) WithAppID(appID string) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithUsername username获取 用户名
func (obj *_RefreshTokenTblMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithExpires expires获取 过期时间
func (obj *_RefreshTokenTblMgr) WithExpires(expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = expires })
}

// GetByOption 功能选项模式获取
func (obj *_RefreshTokenTblMgr) GetByOption(opts ...Option) (result RefreshTokenTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RefreshTokenTblMgr) GetByOptions(opts ...Option) (results []*RefreshTokenTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_RefreshTokenTblMgr) GetFromID(id int) (result RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_RefreshTokenTblMgr) GetBatchFromID(ids []int) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRefreshToken 通过refresh_token获取内容 刷新令牌
func (obj *_RefreshTokenTblMgr) GetFromRefreshToken(refreshToken string) (result RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`refresh_token` = ?", refreshToken).Find(&result).Error

	return
}

// GetBatchFromRefreshToken 批量查找 刷新令牌
func (obj *_RefreshTokenTblMgr) GetBatchFromRefreshToken(refreshTokens []string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`refresh_token` IN (?)", refreshTokens).Find(&results).Error

	return
}

// GetFromTokenType 通过token_type获取内容 令牌类型
func (obj *_RefreshTokenTblMgr) GetFromTokenType(tokenType string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`token_type` = ?", tokenType).Find(&results).Error

	return
}

// GetBatchFromTokenType 批量查找 令牌类型
func (obj *_RefreshTokenTblMgr) GetBatchFromTokenType(tokenTypes []string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`token_type` IN (?)", tokenTypes).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 应用的唯一标识
func (obj *_RefreshTokenTblMgr) GetFromAppID(appID string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`app_id` = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量查找 应用的唯一标识
func (obj *_RefreshTokenTblMgr) GetBatchFromAppID(appIDs []string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`app_id` IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_RefreshTokenTblMgr) GetFromUsername(username string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_RefreshTokenTblMgr) GetBatchFromUsername(usernames []string) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容 过期时间
func (obj *_RefreshTokenTblMgr) GetFromExpires(expires time.Time) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`expires` = ?", expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量查找 过期时间
func (obj *_RefreshTokenTblMgr) GetBatchFromExpires(expiress []time.Time) (results []*RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`expires` IN (?)", expiress).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RefreshTokenTblMgr) FetchByPrimaryKey(id int) (result RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByRefreshToken primary or index 获取唯一内容
func (obj *_RefreshTokenTblMgr) FetchUniqueByRefreshToken(refreshToken string) (result RefreshTokenTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RefreshTokenTbl{}).Where("`refresh_token` = ?", refreshToken).Find(&result).Error

	return
}
