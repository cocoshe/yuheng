package company

import (
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// TODO: 公司进行申诉
func Appeal(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		log.Print(err)
	}

	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := utils.ParseToken(tokenString)
	if err != nil {
		log.Print(err)
		return
	}

	userId := claims.UserId

	nowTime := time.Now().Unix()

	filePath := fmt.Sprintf("static/company/%s_%s", userId, nowTime)
	err = utils.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Print(err)
	}

	//TODO: 更新db
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success to save upload photo",
	})

}
