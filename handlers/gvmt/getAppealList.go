package gvmt

import (
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: 返回所有的企业申诉信息给政府视图(wait for test)
func GetAppealList(c *gin.Context) {
	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}
	var apls []models.Appeal
	drivers.MysqlDb.Table("appeal").Find(&apls)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": apls,
		"msg":  "success",
	})

}