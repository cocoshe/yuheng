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
// @Param       object body models.SelfCheckReq true "自主分析"
// @Success 200 object models.RunModelResp "运行结果"
// @Router       /selfcheck [post]
func SelfCheckHandler(c *gin.Context) {
	ip := config.GlobalConfig.GetString("db.ip")
	r := fmt.Sprintf("http://%s:5001/selfcheck", ip)
	//log.Println("redirect: ", r)
	c.Redirect(http.StatusPermanentRedirect, r)
	//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:5001/run")
}
