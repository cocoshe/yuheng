package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}
