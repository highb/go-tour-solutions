package main

import (
	"fmt"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot take the square root of a negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorNegativeSqrt(x)
	}

	var n int = 10
	var z float64 = 1.0
	var zd float64 = 0.0

	for i := 0; i < n; i++ {
	  zd = (z*z - x) / (2*z)
	  z = z - zd
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
