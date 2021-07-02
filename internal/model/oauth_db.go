package model

import (
	"gorm.io/gorm"
	"time"
)

// AccessTokenTbl 访问令牌
type AccessTokenTbl struct {
	ID          int       `json:"id"`
	AccessToken string    `json:"accessToken"` // 访问令牌
	TokenType   string    `json:"tokenType"`   // 令牌类型
	AppID       string    `json:"appId"`       // 应用的唯一标识
	Username    string    `json:"username"`    // 用户名
	Expires     time.Time `json:"expires"`     // 过期时间
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
	AppID       string
	Username    string
	Expires     string
}{
	ID:          "id",
	AccessToken: "access_token",
	TokenType:   "token_type",
	AppID:       "app_id",
	Username:    "username",
	Expires:     "expires",
}

// Oauth2Tbl oauth2 配置
type Oauth2Tbl struct {
	gorm.Model
	AppID           string    `json:"appId"`           // 应用的唯一标识
	AppKey          string    `json:"appKey"`          // 公匙
	AppSecret       string    `json:"appSecret"`       // 私匙
	Username        string    `json:"username"`        // 用户账号
	ExpireTime      time.Time `json:"expireTime"`      // appid超时时间
	TokenExpireTime int       `json:"tokenExpireTime"` // token过期时间
	StrictSign      bool      `json:"strictSign"`      // 是否强制验签:0：用户自定义，1：强制
	OauthInfo       string    `json:"oauthInfo"`       // oauth信息(一般base64 json串)
	CreatedBy       string    `json:"createdBy"`       // 创建者
	UpdatedBy       string    `json:"updatedBy"`       // 更新者
	DeletedBy       string    `json:"deletedBy"`       // 删除者
}

// TableName get sql table name.获取数据库表名
func (m *Oauth2Tbl) TableName() string {
	return "oauth2_tbl"
}

// Oauth2TblColumns get sql column name.获取数据库列名
var Oauth2TblColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	AppID           string
	AppKey          string
	AppSecret       string
	Username        string
	ExpireTime      string
	TokenExpireTime string
	StrictSign      string
	OauthInfo       string
	CreatedBy       string
	UpdatedBy       string
	DeletedBy       string
}{
	ID:              "id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	AppID:           "app_id",
	AppKey:          "app_key",
	AppSecret:       "app_secret",
	Username:        "username",
	ExpireTime:      "expire_time",
	TokenExpireTime: "token_expire_time",
	StrictSign:      "strict_sign",
	OauthInfo:       "oauth_info",
	CreatedBy:       "created_by",
	UpdatedBy:       "updated_by",
	DeletedBy:       "deleted_by",
}

// RefreshTokenTbl 刷新令牌
type RefreshTokenTbl struct {
	ID           int       `json:"id"`
	RefreshToken string    `json:"refreshToken"` // 刷新令牌
	TokenType    string    `json:"tokenType"`    // 令牌类型
	AppID        string    `json:"appId"`        // 应用的唯一标识
	Username     string    `json:"username"`     // 用户名
	Expires      time.Time `json:"expires"`      // 过期时间
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
	AppID        string
	Username     string
	Expires      string
}{
	ID:           "id",
	RefreshToken: "refresh_token",
	TokenType:    "token_type",
	AppID:        "app_id",
	Username:     "username",
	Expires:      "expires",
}

// UserAccountTbl [...]
type UserAccountTbl struct {
	ID          int         `json:"id"`
	Username    string      `json:"username"`                                                     // 用户账号
	Password    string      `json:"password"`                                                     // 用户密码
	AccountType int         `json:"accountType"`                                                  // 帐号类型:0手机号，1邮件
	AppID       string      `json:"appId"`                                                        // oauth2_tbl表的id(验签id)
	Oauth2Tbl   Oauth2Tbl   `gorm:"joinForeignKey:app_id;foreignKey:app_id" json:"oauth2TblList"` // oauth2 配置
	UserInfoID  int         `json:"userInfoId"`                                                   // 用户附加信息id
	UserInfoTbl UserInfoTbl `gorm:"joinForeignKey:user_info_id;foreignKey:id" json:"userInfoTblList"`
	RegTime     time.Time   `json:"regTime"`   // 注册时间
	RegIP       string      `json:"regIp"`     // 注册ip
	Describ     string      `json:"describ"`   // 描述
	Vaild       bool        `json:"vaild"`     // 是否有效
	CreatedBy   string      `json:"createdBy"` // 创建者
	CreatedAt   time.Time   `json:"createdAt"` // 创建时间
	UpdatedBy   string      `json:"updatedBy"` // 更新者
	UpdatedAt   time.Time   `json:"updatedAt"` // 更新时间
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
	AppID       string
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
	AppID:       "app_id",
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
	ID        int       `json:"id"`
	Username  string    `json:"username"`  // 用户账号
	UserInfo  string    `json:"userInfo"`  // 用户信息(一般base64 json串)
	CreatedBy string    `json:"createdBy"` // 创建者
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedBy string    `json:"updatedBy"` // 更新者
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserInfoTbl) TableName() string {
	return "user_info_tbl"
}

// UserInfoTblColumns get sql column name.获取数据库列名
var UserInfoTblColumns = struct {
	ID        string
	Username  string
	UserInfo  string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	Username:  "username",
	UserInfo:  "user_info",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}
