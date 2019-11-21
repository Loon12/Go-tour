package main

import "fmt"

func main() {
	ch := make(chan int, 200)
	ch <- 1
	ch <- 5
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
