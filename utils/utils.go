package utils

import (
	"math/rand"
	"time"
)

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// newSequence generates random string of 256 bytes
func NewSequence() string {
	rand.Seed(time.Now().UnixNano())

	bs := make([]byte, 127)
	for i := range bs {
		bs[i] = chars[rand.Intn(len(chars))]
	}

	return string(bs)
}

// newSequenceWithLength generates random string of N bytes
func NewSequenceWithLength(n int) string {
	rand.Seed(time.Now().UnixNano())

	bs := make([]byte, n)
	for i := range bs {
		bs[i] = chars[rand.Intn(len(chars))]
	}

	return string(bs)
}
