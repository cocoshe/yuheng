package handlers

import (
	"backend/drivers"
	"github.com/gin-gonic/gin"
)

// @Summary      拿到所有的特征名
// @Description  拿到所有的特征名
// @Accept       json
// @Produce      json
// @Success 200  features and threshold
// @Router       /getFeatures [get]
func GetFeaturesHandler(c *gin.Context) {
	var threshold_body []struct {
		Features  string  `json:"features"`
		Threshold float64 `json:"threshold"`
	}
	drivers.MysqlDb.Table("threshold").Find(&threshold_body)
	thresholdDict := make(map[string]float64)

	for _, v := range threshold_body {
		thresholdDict[v.Features] = v.Threshold
	}
	c.JSON(200, thresholdDict)
}
