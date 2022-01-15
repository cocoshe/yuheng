package visitor

import (
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TODO: 展示提交的历史记录试图  (查一遍表格,看有没有userId一样的accusation,展示积分)(wait for test)
func History(c *gin.Context) {
	var accus []models.Accusation
	var user models.User

	// parsetoken
	tokenString := c.GetHeader("Authorization")
	log.Print(tokenString)
	tokenString = tokenString[7:]
	_, claims, err := utils.ParseToken(tokenString)
	if err != nil {
		log.Print(err)
		return
	}

	userId := claims.UserId
	log.Print("userId: ", userId)
	drivers.MysqlDb.Table("accusation").Where("user_id = ?", userId).Find(&accus)
	drivers.MysqlDb.Table("user").Where("id = ?", userId).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"data":   accus,
		"reward": user.Reward,
		"msg":    "get all hostory for visitor",
	})

}
