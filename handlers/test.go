package handlers

import (
	"backend/dao"
	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	dao.DeleteUserById("20041330")

}
