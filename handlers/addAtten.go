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
)

func AddAttenHandler(c *gin.Context) {
	claim := utils.MiddlewareFunc(c)
	if claim == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token invalid",
		})
		return
	}
	userId := claim.UserId
	var company models.Company
	var user models.User
	err := c.BindJSON(&company)
	if err != nil {
		log.Println(err)
		return
	}
	err = dao.FindUserById(userId, &user)
	if err != nil {
		log.Print(err)
		return
	}
	if user.Atten == "" {
		user.Atten += company.Id
	} else {
		user.Atten += fmt.Sprint(",", company.Id)
	}

	drivers.MysqlDb.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
