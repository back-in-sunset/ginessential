package utils

import "github.com/google/uuid"

// MustUUID returns a uuid string.
func MustUUID() string {
	return uuid.New().String()
}

// NanoID ..
func NanoID() {

}
