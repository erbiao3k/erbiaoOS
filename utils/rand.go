package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

var RandomString = RandLow()

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomLetter(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}

func RandLow() string {
	rand.Seed(time.Now().UnixNano())
	m := md5.New()
	io.WriteString(m, randomLetter(33))
	return hex.EncodeToString(m.Sum(nil))
}
