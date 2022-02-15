package handlers

import (
	"backend/dao"
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// @Summary      获得关注列表的公司id和状态
// @Description  获得关注列表的公司id和状态
// @Accept       json
// @Produce      json
// @Success 200 object models.CpnListResponse
// @Failure 401 object models.FailureResponse
//@Router       /attenInfo [get]
func AttenInfoHandler(c *gin.Context) {
	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}
	userId := claim.UserId
	var user models.User

	err := dao.FindUserById(userId, &user)
	if err != nil {
		log.Print(err)
		return
	}
	attens := user.Atten
	attenList := strings.Split(attens, ",")

	log.Print(attenList)
	var cpnList []models.Company
	for i := 0; i < len(attenList); i++ {
		var t models.Company
		drivers.MysqlDb.Where("company_id = ?", attenList[i]).Find(&t)
		cpnList = append(cpnList, t)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": cpnList,
		"msg":  "success",
	})

}
