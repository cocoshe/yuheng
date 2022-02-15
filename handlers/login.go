package handlers

import (
	"backend/dao"
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary      登录
// @Description  登录
// @Accept       json
// @Produce      json
// @Param         object body models.LoginRequest true "登录信息"
// @Success 200 object models.LoginResponse
// @Failure 401 object models.FailureResponse
//@Router       /login [post]
func LoginHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)

	fmt.Print(user)
	if err != nil {
		log.Print("LoginHandler failed!")
		return
	}
	existUser := &models.User{}
	err = dao.FindUserById(user.Id, existUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "can't find user by id",
		})
		return
	}
	if existUser.Pwd != user.Pwd {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "password wrong",
		})
		return
	}

	token, err := utils.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "server panic",
		})
		log.Print("token generate error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
		"msg":   "success",
	})

}
