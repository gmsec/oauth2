package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1612774359)
	ginrpc.AddGenOne("Oauth.Authorize", "oauth.authorize", []string{"post"})
	ginrpc.AddGenOne("Oauth.CheckToken", "oauth.check_token", []string{"post"})
	ginrpc.AddGenOne("Oauth.Login", "oauth.login", []string{"post"})
	ginrpc.AddGenOne("Oauth.RefreshToken", "oauth.refresh_token", []string{"post"})
}
