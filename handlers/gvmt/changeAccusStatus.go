package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ChangeAccusStatus(c *gin.Context) {
	postDto := &models.AccusStatusDto{}
	err := c.BindJSON(postDto)
	if err != nil {
		log.Println(err)
		return
	}

	post := &models.Accusation{}
	drivers.MysqlDb.Table("accusation").Where("id = ?", postDto.Id).First(post)
	post.Status = postDto.Status
	drivers.MysqlDb.Table("accusation").Save(post)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
