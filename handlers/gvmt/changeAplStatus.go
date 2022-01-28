package gvmt

import (
	"backend/drivers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ChangeAplStatus(c *gin.Context) {
	postDto := &models.AplStatusDto{}
	err := c.BindJSON(postDto)
	if err != nil {
		log.Println(err)
		return
	}

	post := &models.Appeal{}
	drivers.MysqlDb.Table("appeal").Where("id = ?", postDto.Id).First(post)
	post.Status = postDto.Status
	drivers.MysqlDb.Table("appeal").Save(post)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
