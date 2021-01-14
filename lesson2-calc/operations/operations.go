package operations

import (
	"fmt"
	"math"
)

func Sum(a, b float64) (float64, error) {
	return a + b, nil
}

func Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("На 0 делить нельзя.")
	}
	return a / b, nil
}

func Elevate(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}