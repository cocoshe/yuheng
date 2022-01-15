package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
)

// TODO: 政府修改可疑企业状态(wait for test)
func ChangeStatus(c *gin.Context) {
	var apl models.Appeal
	c.ShouldBind(&apl)

	var db_apl models.Appeal
	drivers.MysqlDb.First(&db_apl)

	db_apl.Status = apl.Status
	drivers.MysqlDb.Save(&db_apl)

}
