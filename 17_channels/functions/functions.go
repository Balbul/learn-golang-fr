package functions

import "math/rand"

func GenerateRandomNumber(n int) int {
	return rand.Intn(n)
}
