package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 1000)
	for i := range a {
		a[i] = rand.Intn(100)
	}
	//m := make(map[int]int)
	//printSlice(a)
}
func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}
