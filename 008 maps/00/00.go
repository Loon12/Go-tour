package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func difference2(a1 *[]int, b1 *[]int) {
	a := *a1
	b := *b1
	sort.Ints(a)
	sort.Ints(b)
	fmt.Println(a)
	fmt.Println(b)
	i, j := 0, 0
	for (i < len(a)) && (j < len(b)) {
		if a[i] == b[j] {
			a = append(a[:i], a[i+1:]...)
		} else {
			if a[i] < b[j] {
				i++
			}
			if a[i] > b[j] {
				j++
			}
		}
	}
	(*a1) = a
}

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
	difference2(&a, &b)
	fmt.Println(a)
}
