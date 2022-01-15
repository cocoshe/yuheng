package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: 返回所有的企业申诉信息给政府视图(wait for test)
func Appeal(c *gin.Context) {
	var apls []models.Appeal
	drivers.MysqlDb.Table("appeal").Find(&apls)

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": apls,
		"msg":  "get all appeals for gvmt",
	})

}
