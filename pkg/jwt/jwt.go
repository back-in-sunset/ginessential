package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretDefault = []byte(`key`)

// Claims jwt-claims
type Claims struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userName string, userID string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		UserName: userName,
		UserID:   userID,
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

// FreshToken 刷新token
func FreshToken() {

}
