package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwtSecretDefault = `key`
)

// Claims jwt-claims
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecretDefault)
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecretDefault, nil
	})
	if err != nil {
		return nil, err
	} else if tokenClaims == nil {
		return nil, errors.New("tokenClaims is empty")
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return nil, errors.New("tokenClaims.Claims is error")
	}
	return claims, nil
}
