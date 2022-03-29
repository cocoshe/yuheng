package company

import (
	"backend/drivers"
	"backend/middleware"
	"backend/models"
	"backend/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// @Summary      公司进行申诉
// @Description  公司进行申诉
// @Tags         company
// @Accept       json
// @Produce      json
// @Param         object body models.CompanyAppealDto true "公司申诉信息"
// @Success 200 object models.SuccessResponse
// @Router       /company/appeal [post]
// @Security ApiKeyAuth
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

	claims := utils.MiddlewareFunc(c)
	userId := claims.UserId

	apl := &models.Appeal{
		Id:        uuid.New(),
		CompanyId: cpnId,
		Time:      time.Now(),
		UserId:    userId,
		Post:      post,
		Status:    middleware.WAITING,
	}
	//log.Println(pic)
	if pic != "" {
		bs64 := pic[strings.IndexByte(pic, ',')+1:]
		decodePic, err := base64.StdEncoding.DecodeString(bs64)

		if err != nil {
			log.Println(err)
			return
		}
		err = ioutil.WriteFile("appeal_img/"+apl.Id.String()+".jpg", decodePic, 0644)
		if err != nil {
			return
		}
	}

	drivers.MysqlDb.Table("appeal").Create(apl)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
