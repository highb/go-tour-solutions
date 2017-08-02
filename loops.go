package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var n int = 10
	var z float64 = 1.0
	var zd float64 = 0.0

	for i := 0; i < n; i++ {
	  zd = (z*z - x) / (2*z)
	  z = z - zd
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
