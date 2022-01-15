package visitor

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

// TODO: 游客上传举报信息(wait for test)
func Upload(c *gin.Context) {
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

	filePath := fmt.Sprintf("static/visitor/%s_%s", userId, nowTime)
	err = utils.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Print(err)
	}

	// 更新db
	var accus models.Appeal
	accus.Pic = filePath
	accus.Post = c.PostForm("post")
	accus.Id = uuid.New()
	accus.UserId = userId
	accus.Time = time.Now()
	drivers.MysqlDb.Table("accusation").Create(&accus)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success to create accusation",
	})
}
