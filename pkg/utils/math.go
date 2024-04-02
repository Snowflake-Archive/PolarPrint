package utils

import "math/rand/v2"

func RandomRange(min, max int) int64 {
	return rand.Int64N(int64(max)-int64(min)) + int64(min)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomAlphanumberic(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
