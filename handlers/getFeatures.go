package handlers

import (
	"backend/drivers"
	"github.com/gin-gonic/gin"
)

// @Summary      拿到所有的特征名
// @Description  拿到所有的特征名
// @Accept       json
// @Produce      json
// @Param         object body models.ChangeThresholdReq
// @Success 200 object models.GetFeaturesResp
// @Router       /getFeatures [get]
func GetFeaturesHandler(c *gin.Context) {
	var features []string
	drivers.MysqlDb.Table("threshold").Find(&features)
	c.JSON(200, gin.H{
		"features": features,
	})
}
