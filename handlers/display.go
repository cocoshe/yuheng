package handlers

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary      展示可疑公司及其状态(传入两个索引)
// @Description  展示可疑公司及其状态(传入两个索引)
// @Accept       json
// @Produce      json
// @Param         object body models.DisplayRequest true "传入两个索引"
// @Success 200 object models.DisplayResponse
// @Failure 401 object models.FailureResponse
//@Router       /displayBad [post]
func DisplayHandler(c *gin.Context) {

	var company_infos []models.Company

	var company models.DisplayRequest
	err := c.BindJSON(&company)
	if err != nil {
		log.Println(err)
		return
	}

	idx1 := company.Idx1
	idx2 := company.Idx2
	if idx1 == "" || idx2 == "" {
		drivers.MysqlDb.Table("company").Find(&company_infos)
	} else {
		num1, _ := strconv.Atoi(idx1)
		num2, _ := strconv.Atoi(idx2)
		drivers.MysqlDb.Table("company").Offset(num1 - 1).Limit(num2 - num1 + 1).Find(&company_infos)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": company_infos,
		"msg":  "success",
	})
}
