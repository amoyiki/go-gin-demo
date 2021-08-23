package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	return uuid.NewString()
}
func GenUUID32() string {
	uuidStr := uuid.NewString()
	return strings.ReplaceAll(uuidStr, "-", "")
}
