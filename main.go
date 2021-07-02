package main

import (
	"os"
	"time"

	"oauth2/internal/config"
	"oauth2/internal/routers"
	"oauth2/internal/service/timecallback"

	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	"github.com/gmsec/goplugins/plugin"
	"github.com/gmsec/goplugins/registry/etcdv3"
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
	reg := etcdv3.NewEtcdv3NamingRegistry(clientv3.Config{
		Endpoints:   config.GetEtcdInfo().Addrs,
		DialTimeout: time.Second * time.Duration(config.GetEtcdInfo().Timeout),
	})
	// grpc 相关 初始化服务
	service := micro.NewService(
		micro.WithName("haihuman.srv.oauth2"),
		// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
		micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
		micro.WithRegistryNaming(reg),
	)
	// routers.OnInitService()
	// ----------- end

	// gin restful 相关
	router := gin.Default()
	//router.Use(routers.Cors())
	v1 := router.Group("/oauth2/api/v1")
	routers.OnInitRoot(service.Server(), v1) // 自定义初始化
	//routers.OnInitRouter(v1) // 自定义初始化
	// ------ end

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
