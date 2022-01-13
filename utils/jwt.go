package utils

import (
	"backend/models"
	"github.com/dgrijalva/jwt-go"
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
