package handlers

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary      展示用户排名
// @Description  展示用户排名
// @Accept       json
// @Produce      json
// @Param         object body models.DisplayRequest true "传入两个索引"
// @Success 200 object models.DisplayResponse
// @Failure 401 object models.FailureResponse
// @Router       /rank [post]
func DisplayRank(c *gin.Context) {

	//var company_infos []models.Company
	var user_infos []models.User

	var idx_body models.DisplayRequest
	err := c.BindJSON(&idx_body)
	if err != nil {
		log.Println(err)
		return
	}

	idx1 := idx_body.Idx1
	idx2 := idx_body.Idx2
	if idx1 == "" || idx2 == "" {
		drivers.MysqlDb.Table("user").Order("reward desc").Find(&user_infos)
	} else {
		num1, _ := strconv.Atoi(idx1)
		num2, _ := strconv.Atoi(idx2)
		drivers.MysqlDb.Table("user").Order("reward desc").Offset(num1 - 1).Limit(num2 - num1 + 1).Find(&user_infos)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user_infos,
		"msg":  "success",
	})
}
