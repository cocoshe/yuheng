package utils

import (
	"backend/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("yuheng_secret")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//expirationTime := time.Now().Add(10 * time.Second)
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cocoshe",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}

func MiddlewareFunc(c *gin.Context) *Claims {
	// 获取 authorization header
	tokenString := c.GetHeader("Authorization")

	fmt.Print("请求token: ", tokenString)

	//validate token formate
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "request token error",
		})
		c.Abort()
		return nil
	}

	tokenString = tokenString[7:] //截取字符

	token, claims, err := ParseToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token valid",
		})
		c.Abort()
		return nil
	}

	return claims
}
