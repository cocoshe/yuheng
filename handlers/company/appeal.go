package company

import (
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

// TODO: 公司进行申诉(wait for test)
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

	// 更新db
	var apl models.Appeal
	apl.Pic = filePath
	apl.Post = c.PostForm("post")
	apl.Id = uuid.New()
	apl.UserId = userId
	apl.Time = time.Now()
	drivers.MysqlDb.Table("appeal").Create(&apl)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success to create appeal",
	})

}
