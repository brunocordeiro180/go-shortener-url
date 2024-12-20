package utils

import (
	"math/rand"
)

func RandomId(length int) string {
	randomString := make([]byte, length)

	for i := 0; i < length; i++ {
		opt1 := 65 + rand.Intn(25) //random Upper letter
		opt2 := 97 + rand.Intn(25) //random Lower letter
		choices := []int{opt1, opt2}
		randomString[i] = byte(choices[rand.Intn(len(choices))]) //randomly chose upper or lower
	}

	return string(randomString)
}
