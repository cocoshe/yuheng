package company

import (
	"backend/drivers"
	"backend/middleware"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

// TODO: 公司进行申诉(wait for test)
func Appeal(c *gin.Context) {
	var cpnAplDto models.CompanyAppealDto
	err := c.BindJSON(&cpnAplDto)
	if err != nil {
		log.Println(err)
		return
	}

	cpnId := cpnAplDto.Id
	post := cpnAplDto.Post
	pic := cpnAplDto.Pic

	cpn := &models.Company{}
	drivers.MysqlDb.Table("company").Find(cpn)

	apl := &models.Appeal{
		Id:        uuid.New(),
		CompanyId: cpnId,
		Time:      time.Now(),
		Post:      post,
		Pic:       pic,
		Status:    middleware.WAITING,
	}

	drivers.MysqlDb.Table("appeal").Create(apl)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
