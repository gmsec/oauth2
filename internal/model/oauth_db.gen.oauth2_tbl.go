package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _Oauth2TblMgr struct {
	*_BaseMgr
}

// Oauth2TblMgr open func
func Oauth2TblMgr(db *gorm.DB) *_Oauth2TblMgr {
	if db == nil {
		panic(fmt.Errorf("Oauth2TblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Oauth2TblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("oauth2_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Oauth2TblMgr) GetTableName() string {
	return "oauth2_tbl"
}

// Get 获取
func (obj *_Oauth2TblMgr) Get() (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Oauth2TblMgr) Gets() (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_Oauth2TblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppID app_id获取 应用的唯一标识
func (obj *_Oauth2TblMgr) WithAppID(appID string) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithAppKey app_key获取 公匙
func (obj *_Oauth2TblMgr) WithAppKey(appKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = appKey })
}

// WithAppSecret app_secret获取 私匙
func (obj *_Oauth2TblMgr) WithAppSecret(appSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = appSecret })
}

// WithExpireTime expire_time获取 appid超时时间
func (obj *_Oauth2TblMgr) WithExpireTime(expireTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expire_time"] = expireTime })
}

// WithTokenExpireTime token_expire_time获取 token过期时间
func (obj *_Oauth2TblMgr) WithTokenExpireTime(tokenExpireTime int) Option {
	return optionFunc(func(o *options) { o.query["token_expire_time"] = tokenExpireTime })
}

// WithStrictSign strict_sign获取 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2TblMgr) WithStrictSign(strictSign bool) Option {
	return optionFunc(func(o *options) { o.query["strict_sign"] = strictSign })
}

// GetByOption 功能选项模式获取
func (obj *_Oauth2TblMgr) GetByOption(opts ...Option) (result Oauth2Tbl, err error) {
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
func (obj *_Oauth2TblMgr) GetByOptions(opts ...Option) (results []*Oauth2Tbl, err error) {
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
func (obj *_Oauth2TblMgr) GetFromID(id int) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_Oauth2TblMgr) GetBatchFromID(ids []int) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 应用的唯一标识
func (obj *_Oauth2TblMgr) GetFromAppID(appID string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_id = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量唯一主键查找 应用的唯一标识
func (obj *_Oauth2TblMgr) GetBatchFromAppID(appIDs []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_id IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容 公匙
func (obj *_Oauth2TblMgr) GetFromAppKey(appKey string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key = ?", appKey).Find(&result).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找 公匙
func (obj *_Oauth2TblMgr) GetBatchFromAppKey(appKeys []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key IN (?)", appKeys).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 私匙
func (obj *_Oauth2TblMgr) GetFromAppSecret(appSecret string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_secret = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量唯一主键查找 私匙
func (obj *_Oauth2TblMgr) GetBatchFromAppSecret(appSecrets []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_secret IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容 appid超时时间
func (obj *_Oauth2TblMgr) GetFromExpireTime(expireTime time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expire_time = ?", expireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量唯一主键查找 appid超时时间
func (obj *_Oauth2TblMgr) GetBatchFromExpireTime(expireTimes []time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expire_time IN (?)", expireTimes).Find(&results).Error

	return
}

// GetFromTokenExpireTime 通过token_expire_time获取内容 token过期时间
func (obj *_Oauth2TblMgr) GetFromTokenExpireTime(tokenExpireTime int) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_expire_time = ?", tokenExpireTime).Find(&results).Error

	return
}

// GetBatchFromTokenExpireTime 批量唯一主键查找 token过期时间
func (obj *_Oauth2TblMgr) GetBatchFromTokenExpireTime(tokenExpireTimes []int) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_expire_time IN (?)", tokenExpireTimes).Find(&results).Error

	return
}

// GetFromStrictSign 通过strict_sign获取内容 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2TblMgr) GetFromStrictSign(strictSign bool) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("strict_sign = ?", strictSign).Find(&results).Error

	return
}

// GetBatchFromStrictSign 批量唯一主键查找 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2TblMgr) GetBatchFromStrictSign(strictSigns []bool) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("strict_sign IN (?)", strictSigns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_Oauth2TblMgr) FetchByPrimaryKey(id int) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByAppKey primay or index 获取唯一内容
func (obj *_Oauth2TblMgr) FetchUniqueByAppKey(appKey string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key = ?", appKey).Find(&result).Error

	return
}
