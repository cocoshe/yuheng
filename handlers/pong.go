package handlers

import "github.com/gin-gonic/gin"

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":        "pong",
		"User-Agent": c.GetHeader("User-Agent"),
	})
}
