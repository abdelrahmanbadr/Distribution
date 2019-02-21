package common

import (
	"math"
)

func RoundDivide(num1, num2 int) int {

	return int(math.Round(float64(num1) / float64(num2)))
}
func FloorDivide(num1, num2 int) int {

	return int(math.Floor(float64(num1) / float64(num2)))
}
