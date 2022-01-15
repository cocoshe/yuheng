package middleware

import (
	"backend/dao"
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CompanyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.MiddlewareFunc(c)
		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		var user models.User
		err := dao.FindUserById(userId, &user)
		if err != nil {
			log.Print(err)
			return
		}
		// 验证用户是否存在
		if user.Id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "userId is empty",
			})
			c.Abort()
			return
		}

		if user.Level != COMPANYLEVEL {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "level error",
			})
			c.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
