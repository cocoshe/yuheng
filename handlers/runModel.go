package handlers

import (
	"backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      运行模型, 得到相关数据和结论
// @Description  运行模型, 得到相关数据和结论
// @Accept       json
// @Produce      json
// @Param       object body models.RunModelRequest true "模型运行请求(现在是一些必选项, 后续可以传入自定义的阈值等要求)"
// @Success 200 object models.RunModelResponse "运行结果"
// @Router       /run [post]
func RunHandler(c *gin.Context) {
	ip := config.GlobalConfig.GetString("db.ip")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s:5050/run", ip))
	//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:5000/run")

}
