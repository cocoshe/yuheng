package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      返回所有游客投诉信息列表
// @Description  返回所有游客投诉信息列表
// @Tags         gvmt(admin)
// @Accept       json
// @Produce      json
// @Success 200 object models.AccusationList
// @Router       /gvmt/getAccusationList [get]
// @Security ApiKeyAuth
func GetAccusationList(c *gin.Context) {
	var accus []models.Accusation
	drivers.MysqlDb.Table("accusation").Find(&accus)

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": accus,
		"msg":  "success",
	})
}
