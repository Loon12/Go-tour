package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

const k = 0.0001 // precision

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		for math.Abs(math.Sqrt(x)-z) >= k {
			z = z - ((z*z - x) / (2 * z))
		}
		return z, nil
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
