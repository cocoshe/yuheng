package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      修改某个指标的阈值
// @Description  修改某个指标的阈值
// @Tags         gvmt(admin)
// @Accept       json
// @Produce      json
// @Param       object body models.ChangeThresholdReq true "修改某个指标的阈值"
// @Success 200 object models.SuccessResponse
// @Router       /gvmt/changeThreshold [post]
// @Security ApiKeyAuth
func ChangeThreshold(c *gin.Context) {
	var threshold models.ChangeThresholdReq
	err := c.ShouldBind(&threshold)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if threshold.Features == "" {
		c.JSON(400, gin.H{
			"error": "feature is empty(hint:feature or features?)",
		})
		return
	}
	drivers.MysqlDb.Table("threshold").Where("features = ?", threshold.Features).Update("threshold", threshold.Threshold)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
