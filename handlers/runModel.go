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
// @Param       object body models.RunModelRequest  true "运行模型请求"
// @Success 200 object models.RunModelResp
// @Router       /run [post]
func RunHandler(c *gin.Context) {
	ip := config.GlobalConfig.GetString("db.ip")
	r := fmt.Sprintf("http://%s:5001/run", ip)
	//log.Println("redirect: ", r)
	c.Redirect(http.StatusPermanentRedirect, r)
	//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:5001/run")
}
