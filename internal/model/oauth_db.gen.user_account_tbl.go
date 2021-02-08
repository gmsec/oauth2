package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserAccountTblMgr struct {
	*_BaseMgr
}

// UserAccountTblMgr open func
func UserAccountTblMgr(db *gorm.DB) *_UserAccountTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserAccountTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserAccountTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_account_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserAccountTblMgr) GetTableName() string {
	return "user_account_tbl"
}

// Get 获取
func (obj *_UserAccountTblMgr) Get() (result UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserAccountTblMgr) Gets() (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserAccountTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户账号
func (obj *_UserAccountTblMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 用户密码
func (obj *_UserAccountTblMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithAccountType account_type获取 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) WithAccountType(accountType int) Option {
	return optionFunc(func(o *options) { o.query["account_type"] = accountType })
}

// WithAppKey app_key获取 oauth2_client_tbl表的id(验签id)
func (obj *_UserAccountTblMgr) WithAppKey(appKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = appKey })
}

// WithUserInfoID user_info_id获取 用户附加信息id
func (obj *_UserAccountTblMgr) WithUserInfoID(userInfoID int) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithRegTime reg_time获取 注册时间
func (obj *_UserAccountTblMgr) WithRegTime(regTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["reg_time"] = regTime })
}

// WithRegIP reg_ip获取 注册ip
func (obj *_UserAccountTblMgr) WithRegIP(regIP string) Option {
	return optionFunc(func(o *options) { o.query["reg_ip"] = regIP })
}

// WithDescrib describ获取 描述
func (obj *_UserAccountTblMgr) WithDescrib(describ string) Option {
	return optionFunc(func(o *options) { o.query["describ"] = describ })
}

// WithVaild vaild获取 是否有效
func (obj *_UserAccountTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithCreatedBy created_by获取 创建者
func (obj *_UserAccountTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserAccountTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_UserAccountTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UserAccountTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserAccountTblMgr) GetByOption(opts ...Option) (result UserAccountTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserAccountTblMgr) GetByOptions(opts ...Option) (results []*UserAccountTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserAccountTblMgr) GetFromID(id int) (result UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromID(ids []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUsername 通过username获取内容 用户账号
func (obj *_UserAccountTblMgr) GetFromUsername(username string) (result UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username = ?", username).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUsername 批量唯一主键查找 用户账号
func (obj *_UserAccountTblMgr) GetBatchFromUsername(usernames []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPassword 通过password获取内容 用户密码
func (obj *_UserAccountTblMgr) GetFromPassword(password string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPassword 批量唯一主键查找 用户密码
func (obj *_UserAccountTblMgr) GetBatchFromPassword(passwords []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAccountType 通过account_type获取内容 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) GetFromAccountType(accountType int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account_type = ?", accountType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAccountType 批量唯一主键查找 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) GetBatchFromAccountType(accountTypes []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account_type IN (?)", accountTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAppKey 通过app_key获取内容 oauth2_client_tbl表的id(验签id)
func (obj *_UserAccountTblMgr) GetFromAppKey(appKey string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key = ?", appKey).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAppKey 批量唯一主键查找 oauth2_client_tbl表的id(验签id)
func (obj *_UserAccountTblMgr) GetBatchFromAppKey(appKeys []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key IN (?)", appKeys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserInfoID 通过user_info_id获取内容 用户附加信息id
func (obj *_UserAccountTblMgr) GetFromUserInfoID(userInfoID int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_info_id = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserInfoID 批量唯一主键查找 用户附加信息id
func (obj *_UserAccountTblMgr) GetBatchFromUserInfoID(userInfoIDs []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_info_id IN (?)", userInfoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRegTime 通过reg_time获取内容 注册时间
func (obj *_UserAccountTblMgr) GetFromRegTime(regTime time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reg_time = ?", regTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRegTime 批量唯一主键查找 注册时间
func (obj *_UserAccountTblMgr) GetBatchFromRegTime(regTimes []time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reg_time IN (?)", regTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRegIP 通过reg_ip获取内容 注册ip
func (obj *_UserAccountTblMgr) GetFromRegIP(regIP string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reg_ip = ?", regIP).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRegIP 批量唯一主键查找 注册ip
func (obj *_UserAccountTblMgr) GetBatchFromRegIP(regIPs []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reg_ip IN (?)", regIPs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDescrib 通过describ获取内容 描述
func (obj *_UserAccountTblMgr) GetFromDescrib(describ string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("describ = ?", describ).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDescrib 批量唯一主键查找 描述
func (obj *_UserAccountTblMgr) GetBatchFromDescrib(describs []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("describ IN (?)", describs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVaild 通过vaild获取内容 是否有效
func (obj *_UserAccountTblMgr) GetFromVaild(vaild bool) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild = ?", vaild).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVaild 批量唯一主键查找 是否有效
func (obj *_UserAccountTblMgr) GetBatchFromVaild(vailds []bool) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild IN (?)", vailds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_UserAccountTblMgr) GetFromCreatedBy(createdBy string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_UserAccountTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserAccountTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_UserAccountTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_UserAccountTblMgr) GetFromUpdatedBy(updatedBy string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_UserAccountTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UserAccountTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_UserAccountTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserAccountTblMgr) FetchByPrimaryKey(id int) (result UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByUNIQ5696AD037D3656A4 primay or index 获取唯一内容
func (obj *_UserAccountTblMgr) FetchUniqueByUNIQ5696AD037D3656A4(username string) (result UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username = ?", username).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", result.AppKey).Find(&result.Oauth2Tbl).Error; err != nil { // oauth2 配置
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_info_tbl").Where("id = ?", result.UserInfoID).Find(&result.UserInfoTbl).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByAppKey  获取多个内容
func (obj *_UserAccountTblMgr) FetchIndexByAppKey(appKey string) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("app_key = ?", appKey).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByUserInfoID  获取多个内容
func (obj *_UserAccountTblMgr) FetchIndexByUserInfoID(userInfoID int) (results []*UserAccountTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_info_id = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("oauth2_tbl").Where("app_key = ?", results[i].AppKey).Find(&results[i].Oauth2Tbl).Error; err != nil { // oauth2 配置
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfoTbl).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
