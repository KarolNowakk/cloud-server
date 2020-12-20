package auth

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset generates random string with specified length
//and charset
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//RandomString generates random string with specified length,
//from default charset
func randomString(length int) string {
	return stringWithCharset(length, defaultCharset)
}
