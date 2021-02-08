package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
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
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserInfoTblMgr) Gets() (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserInfoTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNickname nickname获取
func (obj *_UserInfoTblMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithHeadurl headurl获取
func (obj *_UserInfoTblMgr) WithHeadurl(headurl string) Option {
	return optionFunc(func(o *options) { o.query["headurl"] = headurl })
}

// GetByOption 功能选项模式获取
func (obj *_UserInfoTblMgr) GetByOption(opts ...Option) (result UserInfoTbl, err error) {
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
func (obj *_UserInfoTblMgr) GetByOptions(opts ...Option) (results []*UserInfoTbl, err error) {
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
func (obj *_UserInfoTblMgr) GetFromID(id int) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserInfoTblMgr) GetBatchFromID(ids []int) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容
func (obj *_UserInfoTblMgr) GetFromNickname(nickname string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nickname = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找
func (obj *_UserInfoTblMgr) GetBatchFromNickname(nicknames []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nickname IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromHeadurl 通过headurl获取内容
func (obj *_UserInfoTblMgr) GetFromHeadurl(headurl string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("headurl = ?", headurl).Find(&results).Error

	return
}

// GetBatchFromHeadurl 批量唯一主键查找
func (obj *_UserInfoTblMgr) GetBatchFromHeadurl(headurls []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("headurl IN (?)", headurls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserInfoTblMgr) FetchByPrimaryKey(id int) (result UserInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
