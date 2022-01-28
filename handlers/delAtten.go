package handlers

import (
	"backend/dao"
	"backend/drivers"
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func DelAttenHandler(c *gin.Context) {
	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}
	var cpn models.Company
	err := c.BindJSON(&cpn)
	if err != nil {
		log.Println(err)
		return
	}

	var user models.User
	err = dao.FindUserById(claim.UserId, &user)
	if err != nil {
		log.Println(err)
		return
	}

	var updateList []string
	attens := user.Atten
	sepstr := strings.Split(attens, ",")
	for _, str := range sepstr {
		if str != cpn.Id {
			updateList = append(updateList, str)
		}
	}

	update := ""
	if updateList != nil {
		update = updateList[0]
		for i := 1; i < len(updateList); i++ {
			update += fmt.Sprint(",", updateList[i])
		}
	}
	user.Atten = update
	drivers.MysqlDb.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
