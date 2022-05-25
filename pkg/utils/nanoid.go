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
	s, err := gonanoid.Generate("123456789", 36)
	if err != nil {
		panic(err)
	}
	return s
}
