package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Payload struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
}

type CustomClaims struct {
	Payload
	jwt.RegisteredClaims
}

func GenToken(user Payload, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{
		Payload: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func ParseToken(tokenStr string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("token 无效")
	}
	return claims, nil
}
