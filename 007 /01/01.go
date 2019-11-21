package main

import "fmt"

func main() {
	var s []int
	for i := range s {
		s[i]++
	}
	fmt.Println(s)
}
