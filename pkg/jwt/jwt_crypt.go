package jwt

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// Salt 盐
const Salt = `!@)(`

// Scrypt 密码hash
func Scrypt(password, salt string) (string, error) {
	dkpassword, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", dkpassword), nil
}
