package utils

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func GenerateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	var shortCode strings.Builder
	for i := 0; i < 10; i++ {
		shortCode.WriteByte(charset[rand.Intn(len(charset))])
	}
	return shortCode.String()
}
