package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: 返回所有游客投诉信息列表(wait for test)
func GetAccusationList(c *gin.Context) {
	var accus []models.Accusation
	drivers.MysqlDb.Table("appeal").Find(&accus)

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": accus,
		"msg":  "get all accusations for gvmt",
	})
}
