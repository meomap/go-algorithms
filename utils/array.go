// Package utils provides utility for array
package utils

import (
	"math/rand"
	"time"
)

func RandNumbers(numbers uint) (out []uint) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	out = make([]uint, numbers)
	for i := uint(0); i < numbers; i++ {
		out[i] = uint(r1.Intn(10000))
	}
	return
}
