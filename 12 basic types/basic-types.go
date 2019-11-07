package main

import (
	"fmt"
	"math/cmplx"
)

var (
	//ToBe is importent for me
	ToBe bool = false
	//MaxInt is importent for me
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %#v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %#v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %#v\n", z, z)
}
