package main

import (
	"os"
	"time"

	"oauth2/internal/config"
	"oauth2/internal/routers"
	"oauth2/internal/service/timecallback"

	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/micro"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/server"
	"github.com/xxjwxc/public/timerDeal"
)

// CallBack service call backe
func CallBack() {
	//设置时间回调
	var tcb timerDeal.TimerDeal
	tcb.AddOneCall(timecallback.TimeCallBackToken)
	tcb.SetCallBackTimer(24 * time.Hour)
	tcb.OnSart()
	//---------------------end

	// swagger
	myswagger.SetHost("https://localhost:8080")
	myswagger.SetBasePath("oauth2")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("oauth2.srv.eg1"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		// micro.WithRegistryNaming(reg),
	)
	//routers.OnInitService()
	// ----------- end

	// gin restful 相关
	router := gin.Default()
	//router.Use(routers.Cors())
	v1 := router.Group("/oauth2/api/v1")
	//routers.OnInitRouter(v1) // 自定义初始化
	// ------ end

	routers.OnInitRoot(service.Server(), v1)

	plg, b := plugin.Run(plugin.WithMicro(service),
		plugin.WithGin(router),
		plugin.WithAddr(":"+config.GetPort()))

	if b == nil {
		plg.Wait()
	}
	mylog.Info("done")
}

func main() {
	if config.GetIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
