package visitor

import (
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TODO: 展示提交的历史记录试图  (查一遍表格,看有没有userId一样的accusation)(wait for test)
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

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": accus,
		//"reward": user.Reward,
		//"msg":    "get all history for visitor",
	})

}
