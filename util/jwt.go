package util

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// SigningKey 加密key
var SigningKey = []byte("entry-task")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToke 生成Token
func GenerateToke(username string) (string, error) {
	claims := Claims{username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,    // 生效时间，当前时间的一分钟之前生效
			ExpiresAt: time.Now().Unix() + 60*60, // 失效时间，当前时间一小时后失效
			Issuer:    username,                  // 签发者
		},
	}
	// 加密方式 HS256
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SigningKey)
	return token, err
}

// VerifyToken 验证前端传入的Token
func VerifyToken(tokenString string) (*jwt.Token, *Claims, error) {
	chaims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, chaims, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return token, chaims, err
}
