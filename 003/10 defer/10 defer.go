package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println(x + 1)
	defer fmt.Println("DOBRY")
	fmt.Println("Hi")
	fmt.Println(x)
}
