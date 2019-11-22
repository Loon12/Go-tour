//Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 1000)
	b := make([]int, 1000)
	for i := range a {
		a[i] = rand.Intn(100000)
		b[i] = rand.Intn(100000)
	}
	//включал сортировку для удобства проверки результата
	//a = sorT(a)
	//b = sorT(b)
	fmt.Println("\nПервый Массив:")
	printSlice(a)
	fmt.Println("\nВторой Массив:")
	printSlice(b)
	m := make(map[int]int)
	for i := range a {
		if _, th := m[a[i]]; th == false {
			m[a[i]]++
		}
	}
	fmt.Println("\n\nЧисла, которые есть и в первом и во втором массиве:")
	for i := range b {
		if _, th := m[b[i]]; th == true {
			if m[b[i]] == 1 {
				fmt.Printf("%v ", b[i])
				m[b[i]]++
			}
		}
	}

}

func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}

func sorT(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
	return a
}
