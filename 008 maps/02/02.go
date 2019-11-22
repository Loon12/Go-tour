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
	printSlice(a)
	m := make(map[int]int)
	for i:=range a {
		if _,th:=m[a[i]];th==false{
			m[a[i]]++
			fmt.Printf("%v ",a[i])
		}
	}
}
func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}
