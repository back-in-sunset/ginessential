package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// MustNanoID if have error will be panic
func MustNanoID() string {
	return gonanoid.Must()
}

// NanoNumbID ..
func NanoNumbID() string {
	s, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789", 8)
	if err != nil {
		panic(err)
	}
	return s
}
