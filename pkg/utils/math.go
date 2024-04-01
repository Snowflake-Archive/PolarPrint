package utils

import "math/rand/v2"

func RandomRange(min, max int) int64 {
	return rand.Int64N(int64(max)-int64(min)) + int64(min)
}
