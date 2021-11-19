package utils

import (
	"bytes"
	"math/rand"
	"time"
)

const char string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(n int) string {
	var s bytes.Buffer
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < n; i++{
		s.WriteByte(char[rand.Intn(len(char))])
	}
	return s.String()
}
