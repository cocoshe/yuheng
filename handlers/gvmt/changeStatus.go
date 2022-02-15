package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary      政府修改企业状态
// @Description  政府修改企业状态
// @Tags         gvmt(admin)
// @Accept       json
// @Produce      json
// @Param         object body models.Company true "修改公司状态为status, 1为正常, 2为可疑违规"
// @Success 200 object models.SuccessResponse
//@Router       /gvmt/changeStatus [post]
func ChangeStatus(c *gin.Context) {
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
