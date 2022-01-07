package routers

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", handlers.IndexHandler)
	r.GET("/ping", handlers.PingHandler)
}
