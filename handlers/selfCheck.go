package handlers

import (
	"backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      自主分析
// @Description  自主分析
// @Accept       json
// @Produce      json
// @Param       object body models.RunModelRequest true "模型运行请求(现在是一些必选项, 后续可以传入自定义的阈值等要求)"
// @Success 200 object models.RunModelResp "运行结果"
// @Router       /selfcheck [post]
func SelfCheckHandler(c *gin.Context) {
	ip := config.GlobalConfig.GetString("db.ip")
	r := fmt.Sprintf("http://%s:5001/selfcheck", ip)
	//log.Println("redirect: ", r)
	c.Redirect(http.StatusPermanentRedirect, r)
	//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:5001/run")
}
