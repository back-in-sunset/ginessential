package utils

import "github.com/google/uuid"

// NewUUID returns a uuid string.
func NewUUID() string {
	return uuid.New().String()
}

// NanoID ..
func NanoID() {

}
