package main

import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("Hello", "Aleksandr")
	fmt.Println(a, b)
}
