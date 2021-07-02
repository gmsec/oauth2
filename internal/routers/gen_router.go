package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1624614500)
	ginrpc.AddGenOne("Oauth.Authorize", "oauth.authorize", []string{"post"})
	ginrpc.AddGenOne("Oauth.CheckToken", "oauth.check_token", []string{"post"})
	ginrpc.AddGenOne("Oauth.CreateOauth", "oauth.create_oauth", []string{"post"})
	ginrpc.AddGenOne("Oauth.CreateUser", "oauth.create_user", []string{"post"})
	ginrpc.AddGenOne("Oauth.DeleteApp", "oauth.delete_app", []string{"post"})
	ginrpc.AddGenOne("Oauth.GetAppList", "oauth.get_app_list", []string{"post"})
	ginrpc.AddGenOne("Oauth.GetLoginInfo", "oauth.get_login_info", []string{"post"})
	ginrpc.AddGenOne("Oauth.GetUsers", "oauth.get_users", []string{"post"})
	ginrpc.AddGenOne("Oauth.Login", "oauth.login", []string{"post"})
	ginrpc.AddGenOne("Oauth.RefreshToken", "oauth.refresh_token", []string{"post"})
	ginrpc.AddGenOne("Oauth.UpdateUser", "oauth.update_user", []string{"post"})
}
