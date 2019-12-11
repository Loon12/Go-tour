package main

import "fmt"

func ten(s []int) {
	for i := range s {
		s[i] = 10
	}
}

func addone(s []int) []int {
	s = append(s, 1)
	return s
}

func main() {
	a := make([]int, 9)
	ten(a)
	a = addone(a)
	fmt.Println(a)
}
