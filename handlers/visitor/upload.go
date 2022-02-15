package visitor

import (
	"backend/drivers"
	"backend/middleware"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
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
//@Router       /visitor/upload [post]
// TODO: 游客上传举报信息(wait for test)
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
		Pic:       userPostDto.Pic,
		Status:    middleware.WAITING,
	}

	drivers.MysqlDb.Table("accusation").Create(accus)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
