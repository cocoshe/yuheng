package middleware

import (
	"backend/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	JWTTokenMalformed   = errors.New("jwt token malformed")
	JWTTokenExpired     = errors.New("jwt token expired")
	JWTTokenNotValidYet = errors.New("jwt token not valid yet")
	JWTTokenInvalid     = errors.New("jwt token invalid")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(config.GlobalConfig.GetString("basic.jwt.secret")),
	}
}

type JWTClaims struct {
	jwt.StandardClaims
}

// create token
func (j JWT) CreateToken(claims JWTClaims) (string, error) {
	//c := MyClains{
	//	Username: "cocoshe",
	//	StandardClaims: jwt.StandardClaims{
	//		NotBefore: time.Now().Unix() - 60,
	//		ExpiresAt: time.Now().Unix() + 60*60*24,
	//		Issuer:    "cocoshe",
	//	},
	//}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Print(token)
	return token.SignedString(j.SigningKey)

}

// parse token
func (j JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, JWTTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, JWTTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, JWTTokenNotValidYet
			} else {
				return nil, JWTTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, JWTTokenInvalid

}
