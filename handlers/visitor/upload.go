package visitor

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

// @Summary      进行举报
// @Description  进行举报
// @Tags         visitor
// @Accept       json
// @Produce      json
// @Param         object body models.UserPostDto true "举报信息"
// @Success 200 object models.SuccessResponse
// @Failure 401 object models.FailureResponse
// @Router       /visitor/upload [post]
// @Security ApiKeyAuth
func Upload(c *gin.Context) {
	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}

	var userPostDto models.UserPostDto
	err := c.BindJSON(&userPostDto)
	if err != nil {
		log.Println(err)
		return
	}

	accus := &models.Accusation{
		Id:        uuid.New(),
		CompanyId: userPostDto.Id,
		UserId:    claim.UserId,
		Time:      time.Now(),
		Post:      userPostDto.Post,
		Status:    middleware.WAITING,
	}

	pic := userPostDto.Pic
	if pic != "" {
		idx1 := strings.IndexByte(pic, ',')
		bs64 := pic[idx1+1:]
		prefix := pic[:idx1]
		idx2 := strings.IndexByte(prefix, '/')
		idx3 := strings.IndexByte(prefix, ';')
		picType := prefix[idx2+1 : idx3]
		decodePic, err := base64.StdEncoding.DecodeString(bs64)

		if err != nil {
			log.Println(err)
			return
		}
		accus.PicType = picType
		err = ioutil.WriteFile("accus_img/"+accus.Id.String()+"."+picType, decodePic, 0644)
		if err != nil {
			return
		}
	}

	drivers.MysqlDb.Table("accusation").Create(accus)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
