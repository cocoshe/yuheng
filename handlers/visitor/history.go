package visitor

import (
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// @Summary      返回自己的举报记录
// @Description  返回自己的举报记录
// @Tags         visitor
// @Accept       json
// @Produce      json
// @Success 200 object models.AccusationList
// @Failure 401 object models.FailureResponse
// @Router       /visitor/history [get]
// @Security ApiKeyAuth
func History(c *gin.Context) {
	var accus []models.Accusation
	//var user models.User

	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}

	userId := claim.UserId
	log.Print("userId: ", userId)
	drivers.MysqlDb.Table("accusation").Where("user_id = ?", userId).Find(&accus)
	//drivers.MysqlDb.Table("user").Where("id = ?", userId).First(&user)

	for i, acc := range accus {
		tempId := acc.Id
		picPath := "accus_img/" + tempId.String() + "." + acc.PicType
		f, _ := ioutil.ReadFile(picPath)
		accus[i].Pic = "data:image/png;base64," + base64.StdEncoding.EncodeToString(f)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": accus,
		//"reward": user.Reward,
		"msg": "success",
	})

}
