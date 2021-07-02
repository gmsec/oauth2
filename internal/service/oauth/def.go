package oauth

// TokenInfo token信息
type TokenInfo struct {
	AccessToken       string `json:"access_token"`        // 访问令牌
	RefreshToken      string `json:"refresh_token"`       // 刷新令牌
	AccessExpireTime  int64  `json:"access_expire_time"`  // 访问令牌过期时间
	RefreshExpireTime int64  `json:"refresh_expire_time"` // 刷新令牌过期时间
	UserName          string `json:"user_name"`           // 用户名
}

// TokenCache token信息
type TokenCache struct {
	Token      string `json:"token"`              // 令牌
	ExpireTime int64  `json:"access_expire_time"` //过期时间
	UserName   string `json:"user_info"`          // 附加信息
	AppID      string `json:"app_key"`            // 应用key
	TokenType  string `json:"token_type"`         // 授权类型
}
