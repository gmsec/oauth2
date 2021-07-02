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
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Oauth2TblMgr) Gets() (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_Oauth2TblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_Oauth2TblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_Oauth2TblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_Oauth2TblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
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

// WithUsername username获取 用户账号
func (obj *_Oauth2TblMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
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

// WithOauthInfo oauth_info获取 oauth信息(一般base64 json串)
func (obj *_Oauth2TblMgr) WithOauthInfo(oauthInfo string) Option {
	return optionFunc(func(o *options) { o.query["oauth_info"] = oauthInfo })
}

// WithCreatedBy created_by获取 创建者
func (obj *_Oauth2TblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_Oauth2TblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_Oauth2TblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// GetByOption 功能选项模式获取
func (obj *_Oauth2TblMgr) GetByOption(opts ...Option) (result Oauth2Tbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(options.query).Find(&result).Error

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

	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_Oauth2TblMgr) GetFromID(id int64) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary key
func (obj *_Oauth2TblMgr) GetBatchFromID(ids []int64) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_Oauth2TblMgr) GetFromCreatedAt(createdAt time.Time) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量查找 created time
func (obj *_Oauth2TblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_Oauth2TblMgr) GetFromUpdatedAt(updatedAt time.Time) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 updated at
func (obj *_Oauth2TblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_Oauth2TblMgr) GetFromDeletedAt(deletedAt time.Time) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量查找 deleted time
func (obj *_Oauth2TblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where(" IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 应用的唯一标识
func (obj *_Oauth2TblMgr) GetFromAppID(appID string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_id` = ?", appID).Find(&result).Error

	return
}

// GetBatchFromAppID 批量查找 应用的唯一标识
func (obj *_Oauth2TblMgr) GetBatchFromAppID(appIDs []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_id` IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容 公匙
func (obj *_Oauth2TblMgr) GetFromAppKey(appKey string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_key` = ?", appKey).Find(&result).Error

	return
}

// GetBatchFromAppKey 批量查找 公匙
func (obj *_Oauth2TblMgr) GetBatchFromAppKey(appKeys []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_key` IN (?)", appKeys).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 私匙
func (obj *_Oauth2TblMgr) GetFromAppSecret(appSecret string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_secret` = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量查找 私匙
func (obj *_Oauth2TblMgr) GetBatchFromAppSecret(appSecrets []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_secret` IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户账号
func (obj *_Oauth2TblMgr) GetFromUsername(username string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户账号
func (obj *_Oauth2TblMgr) GetBatchFromUsername(usernames []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容 appid超时时间
func (obj *_Oauth2TblMgr) GetFromExpireTime(expireTime time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`expire_time` = ?", expireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量查找 appid超时时间
func (obj *_Oauth2TblMgr) GetBatchFromExpireTime(expireTimes []time.Time) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`expire_time` IN (?)", expireTimes).Find(&results).Error

	return
}

// GetFromTokenExpireTime 通过token_expire_time获取内容 token过期时间
func (obj *_Oauth2TblMgr) GetFromTokenExpireTime(tokenExpireTime int) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`token_expire_time` = ?", tokenExpireTime).Find(&results).Error

	return
}

// GetBatchFromTokenExpireTime 批量查找 token过期时间
func (obj *_Oauth2TblMgr) GetBatchFromTokenExpireTime(tokenExpireTimes []int) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`token_expire_time` IN (?)", tokenExpireTimes).Find(&results).Error

	return
}

// GetFromStrictSign 通过strict_sign获取内容 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2TblMgr) GetFromStrictSign(strictSign bool) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`strict_sign` = ?", strictSign).Find(&results).Error

	return
}

// GetBatchFromStrictSign 批量查找 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2TblMgr) GetBatchFromStrictSign(strictSigns []bool) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`strict_sign` IN (?)", strictSigns).Find(&results).Error

	return
}

// GetFromOauthInfo 通过oauth_info获取内容 oauth信息(一般base64 json串)
func (obj *_Oauth2TblMgr) GetFromOauthInfo(oauthInfo string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`oauth_info` = ?", oauthInfo).Find(&results).Error

	return
}

// GetBatchFromOauthInfo 批量查找 oauth信息(一般base64 json串)
func (obj *_Oauth2TblMgr) GetBatchFromOauthInfo(oauthInfos []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`oauth_info` IN (?)", oauthInfos).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_Oauth2TblMgr) GetFromCreatedBy(createdBy string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_Oauth2TblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_Oauth2TblMgr) GetFromUpdatedBy(updatedBy string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找 更新者
func (obj *_Oauth2TblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_Oauth2TblMgr) GetFromDeletedBy(deletedBy string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`deleted_by` = ?", deletedBy).Find(&results).Error

	return
}

// GetBatchFromDeletedBy 批量查找 删除者
func (obj *_Oauth2TblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`deleted_by` IN (?)", deletedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_Oauth2TblMgr) FetchByPrimaryKey(id int64) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByAppID primary or index 获取唯一内容
func (obj *_Oauth2TblMgr) FetchUniqueByAppID(appID string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_id` = ?", appID).Find(&result).Error

	return
}

// FetchUniqueByAppKey primary or index 获取唯一内容
func (obj *_Oauth2TblMgr) FetchUniqueByAppKey(appKey string) (result Oauth2Tbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Oauth2Tbl{}).Where("`app_key` = ?", appKey).Find(&result).Error

	return
}
