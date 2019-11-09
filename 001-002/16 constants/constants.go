package main

import "fmt"

//Pi is importent for me
const Pi = 3.14

func main() {
	const World = "DOBRY_BORODA"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
