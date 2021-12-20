package util

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomDigit(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func RandomLenDigit(len int) string {
	var result = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
}
