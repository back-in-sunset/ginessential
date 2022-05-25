package jwt

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT ..
type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	nowFn      func() time.Time
	sub        string
}

// TimeFnOption jwt time opthion
func TimeFnOption(nowfn func() time.Time) OptFn {
	return func(j *JWT) {
		j.nowFn = nowfn
	}
}

// SubFnOption sub filed
func SubFnOption(sub string) OptFn {
	return func(j *JWT) {
		j.sub = sub
	}
}

// OptFn ..
type OptFn func(j *JWT)

// NewJWT ..
func NewJWT(privateKey, publicKey []byte, ops ...OptFn) (*JWT, error) {
	prikey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}
	pubkey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	j := &JWT{
		privateKey: prikey,
		publicKey:  pubkey,
	}
	for _, op := range ops {
		op(j)
	}
	return j, nil
}

// Create ..
func (j *JWT) Create(content string, expDuration time.Duration) (string, error) {
	claims := make(jwt.MapClaims)
	now := j.nowFn()
	claims["dat"] = content                     // Our custom data.
	claims["exp"] = now.Add(expDuration).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()                  // The time at which the token was issued.
	claims["sub"] = j.sub

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}
