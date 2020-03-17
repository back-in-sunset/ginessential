package util

import (
	"math/rand"
	"time"
)

// RandomString 随机字符串
func RandomString(n int) string {
	letters := []byte("asdfasGUYGUdsadOAAsfmcsoferafsfzvsgg")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
