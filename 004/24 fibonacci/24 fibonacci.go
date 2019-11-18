package main

import "fmt"

func fibonacci() func() int {
	x0 := 0
	x1 := 1
	return func() int {
		x := x0 + x1
		x0 = x1
		x1 = x
		return x0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 12; i++ {
		fmt.Println(f())
	}
}
