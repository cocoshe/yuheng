package server

import (
	"backend/config"
	"backend/routers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine) {
	//服务配置
	httpServer = gin.Default()
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AddAllowHeaders("Authorization")

	httpServer.Use(cors.New(corsConf))

	httpServer.LoadHTMLGlob(config.GlobalConfig.GetString("httpEngine.templates"))
	//注册路由
	routers.RegisterRoutes(httpServer)

	Addr := config.GlobalConfig.GetString("httpEngine.addr")
	Port := config.GlobalConfig.GetString("httpEngine.port")
	serverAddr := fmt.Sprint(Addr, ":", Port)

	err := httpServer.Run(serverAddr)
	if nil != err {
		panic("server run error: " + err.Error())
	}

}
