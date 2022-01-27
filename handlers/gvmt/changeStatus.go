package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func ChangeStatus_v2(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		log.Println(err)
		return
	}

	var db_company models.Company
	drivers.MysqlDb.Table("company").
		Where("company_id = ?", company.Id).
		First(&db_company)
	db_company.Status = company.Status
	log.Println(db_company)
	log.Println(company)
	drivers.MysqlDb.Save(&db_company)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
