package main

import (
	_ "backend/config"
	"backend/drivers"
	"backend/server"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 服务停止时清除数据库链接
	defer drivers.MysqlDb.Close()

	//启动服务
	server.Run(HttpServer)
}
