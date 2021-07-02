package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserInfoTblMgr struct {
	*_BaseMgr
}

// UserInfoTblMgr open func
func UserInfoTblMgr(db *gorm.DB) *_UserInfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserInfoTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserInfoTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_info_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserInfoTblMgr) GetTableName() string {
	return "user_info_tbl"
}

// Get 获取
func (obj *_UserInfoTblMgr) Get() (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserInfoTblMgr) Gets() (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserInfoTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户账号
func (obj *_UserInfoTblMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithUserInfo user_info获取 用户信息(一般base64 json串)
func (obj *_UserInfoTblMgr) WithUserInfo(userInfo string) Option {
	return optionFunc(func(o *options) { o.query["user_info"] = userInfo })
}

// WithCreatedBy created_by获取 创建者
func (obj *_UserInfoTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserInfoTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_UserInfoTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UserInfoTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserInfoTblMgr) GetByOption(opts ...Option) (result UserInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserInfoTblMgr) GetByOptions(opts ...Option) (results []*UserInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserInfoTblMgr) GetFromID(id int) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserInfoTblMgr) GetBatchFromID(ids []int) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户账号
func (obj *_UserInfoTblMgr) GetFromUsername(username string) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`username` = ?", username).Find(&result).Error

	return
}

// GetBatchFromUsername 批量查找 用户账号
func (obj *_UserInfoTblMgr) GetBatchFromUsername(usernames []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromUserInfo 通过user_info获取内容 用户信息(一般base64 json串)
func (obj *_UserInfoTblMgr) GetFromUserInfo(userInfo string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`user_info` = ?", userInfo).Find(&results).Error

	return
}

// GetBatchFromUserInfo 批量查找 用户信息(一般base64 json串)
func (obj *_UserInfoTblMgr) GetBatchFromUserInfo(userInfos []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`user_info` IN (?)", userInfos).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_UserInfoTblMgr) GetFromCreatedBy(createdBy string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_UserInfoTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserInfoTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_UserInfoTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_UserInfoTblMgr) GetFromUpdatedBy(updatedBy string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找 更新者
func (obj *_UserInfoTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UserInfoTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_UserInfoTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserInfoTblMgr) FetchByPrimaryKey(id int) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUsername primary or index 获取唯一内容
func (obj *_UserInfoTblMgr) FetchUniqueByUsername(username string) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfoTbl{}).Where("`username` = ?", username).Find(&result).Error

	return
}
