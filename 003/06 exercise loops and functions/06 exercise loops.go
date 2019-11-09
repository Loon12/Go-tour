package main

import (
	"fmt"
	"math"
)

const k = 0.0001 // precision

func sqrt(x float64) float64 {
	var z float64 = 1.0
	for math.Abs(math.Sqrt(x)-z) >= k {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(sqrt(2))
}
