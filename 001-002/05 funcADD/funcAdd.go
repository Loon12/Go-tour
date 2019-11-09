package main

import "fmt"

func add1(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add1(12, 13))
	fmt.Println(add2(12, 13))
}
