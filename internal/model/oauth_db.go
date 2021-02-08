package model

import (
	"time"
)

// AccessTokenTbl 访问令牌
type AccessTokenTbl struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	AccessToken string    `gorm:"unique" json:"access_token"` // 访问令牌
	TokenType   string    `json:"token_type"`                 // 令牌类型
	AppKey      string    `json:"app_key"`                    // key
	Userinfo    string    `json:"userinfo"`                   // 用户名
	Expires     time.Time `json:"expires"`                    // 过期时间
}

// TableName get sql table name.获取数据库表名
func (m *AccessTokenTbl) TableName() string {
	return "access_token_tbl"
}

// AccessTokenTblColumns get sql column name.获取数据库列名
var AccessTokenTblColumns = struct {
	ID          string
	AccessToken string
	TokenType   string
	AppKey      string
	Userinfo    string
	Expires     string
}{
	ID:          "id",
	AccessToken: "access_token",
	TokenType:   "token_type",
	AppKey:      "app_key",
	Userinfo:    "userinfo",
	Expires:     "expires",
}

// Oauth2Tbl oauth2 配置
type Oauth2Tbl struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	AppID           string    `json:"app_id"`                       // 应用的唯一标识
	AppKey          string    `gorm:"unique" json:"app_key"`        // 公匙
	AppSecret       string    `json:"app_secret"`                   // 私匙
	ExpireTime      time.Time `json:"expire_time"`                  // appid超时时间
	TokenExpireTime int       `json:"token_expire_time"`            // token过期时间
	StrictSign      bool      `gorm:"default:1" json:"strict_sign"` // 是否强制验签:0：用户自定义，1：强制
}

// TableName get sql table name.获取数据库表名
func (m *Oauth2Tbl) TableName() string {
	return "oauth2_tbl"
}

// Oauth2TblColumns get sql column name.获取数据库列名
var Oauth2TblColumns = struct {
	ID              string
	AppID           string
	AppKey          string
	AppSecret       string
	ExpireTime      string
	TokenExpireTime string
	StrictSign      string
}{
	ID:              "id",
	AppID:           "app_id",
	AppKey:          "app_key",
	AppSecret:       "app_secret",
	ExpireTime:      "expire_time",
	TokenExpireTime: "token_expire_time",
	StrictSign:      "strict_sign",
}

// RefreshTokenTbl 刷新令牌
type RefreshTokenTbl struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	RefreshToken string    `gorm:"unique" json:"refresh_token"` // 刷新令牌
	TokenType    string    `json:"token_type"`                  // 令牌类型
	AppKey       string    `json:"app_key"`                     // 访问令牌
	Userinfo     string    `json:"userinfo"`                    // 用户名
	Expires      time.Time `json:"expires"`                     // 过期时间
}

// TableName get sql table name.获取数据库表名
func (m *RefreshTokenTbl) TableName() string {
	return "refresh_token_tbl"
}

// RefreshTokenTblColumns get sql column name.获取数据库列名
var RefreshTokenTblColumns = struct {
	ID           string
	RefreshToken string
	TokenType    string
	AppKey       string
	Userinfo     string
	Expires      string
}{
	ID:           "id",
	RefreshToken: "refresh_token",
	TokenType:    "token_type",
	AppKey:       "app_key",
	Userinfo:     "userinfo",
	Expires:      "expires",
}

// UserAccountTbl [...]
type UserAccountTbl struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Username    string      `gorm:"unique" json:"username"`                                           // 用户账号
	Password    string      `json:"password"`                                                         // 用户密码
	AccountType int         `gorm:"default:0" json:"account_type"`                                    // 帐号类型:0手机号，1邮件
	AppKey      string      `gorm:"index:app_key;default:nomal" json:"app_key"`                       // oauth2_client_tbl表的id(验签id)
	Oauth2Tbl   Oauth2Tbl   `gorm:"joinForeignKey:app_key;foreignKey:app_key" json:"oauth2_tbl_list"` // oauth2 配置
	UserInfoID  int         `gorm:"index:user_info_id" json:"user_info_id"`                           // 用户附加信息id
	UserInfoTbl UserInfoTbl `gorm:"joinForeignKey:user_info_id;foreignKey:id" json:"user_info_tbl_list"`
	RegTime     time.Time   `json:"reg_time"`                     // 注册时间
	RegIP       string      `json:"reg_ip"`                       // 注册ip
	Describ     string      `json:"describ"`                      // 描述
	Vaild       bool        `gorm:"default:1" json:"vaild"`       // 是否有效
	CreatedBy   string      `gorm:"default:''" json:"created_by"` // 创建者
	CreatedAt   time.Time   `json:"created_at"`                   // 创建时间
	UpdatedBy   string      `gorm:"default:''" json:"updated_by"` // 更新者
	UpdatedAt   time.Time   `json:"updated_at"`                   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserAccountTbl) TableName() string {
	return "user_account_tbl"
}

// UserAccountTblColumns get sql column name.获取数据库列名
var UserAccountTblColumns = struct {
	ID          string
	Username    string
	Password    string
	AccountType string
	AppKey      string
	UserInfoID  string
	RegTime     string
	RegIP       string
	Describ     string
	Vaild       string
	CreatedBy   string
	CreatedAt   string
	UpdatedBy   string
	UpdatedAt   string
}{
	ID:          "id",
	Username:    "username",
	Password:    "password",
	AccountType: "account_type",
	AppKey:      "app_key",
	UserInfoID:  "user_info_id",
	RegTime:     "reg_time",
	RegIP:       "reg_ip",
	Describ:     "describ",
	Vaild:       "vaild",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
}

// UserInfoTbl [...]
type UserInfoTbl struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Nickname string `json:"nickname"`
	Headurl  string `json:"headurl"`
}

// TableName get sql table name.获取数据库表名
func (m *UserInfoTbl) TableName() string {
	return "user_info_tbl"
}

// UserInfoTblColumns get sql column name.获取数据库列名
var UserInfoTblColumns = struct {
	ID       string
	Nickname string
	Headurl  string
}{
	ID:       "id",
	Nickname: "nickname",
	Headurl:  "headurl",
}
