package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := make([]int, 10)
	for i := range a {
		a[i] = rand.Intn(10)
	}

	b := make([]int, 10)
	for i := range b {
		b[i] = rand.Intn(10)
	}

	fmt.Println(a)
	fmt.Println(b)

	delel(&a, &b)

	fmt.Println(a)

}

func delel(a *[]int, b *[]int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len(*b); j++ {
			if (*a)[i] == (*b)[j] {
				if i == len(*a)-1 {
					*a = (*a)[:i]
					break
				} else {
					*a = append((*a)[:i], (*a)[i+1:]...)
					j = -1
				}
			}
		}
	}
}

func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}
