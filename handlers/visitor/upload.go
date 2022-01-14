package visitor

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
)

// TODO: 游客上传举报信息
func Upload(c *gin.Context) {
	var accusation models.Accusation
	err := c.ShouldBind(&accusation)
	if err != nil {
		log.Print("upload accusation failed")
		return
	}
}
