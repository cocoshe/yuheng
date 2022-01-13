package routers

import (
	"backend/handlers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/index", handlers.IndexHandler)
	r.GET("/ping", handlers.PingHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/info", middleware.AuthMiddleware(), handlers.Info)

}
