package main

import (
	"fmt"
)

func main() {
	//var t int
	//fmt.Print("Введите количество элементов массива: ")
	//fmt.Scanf("%d \n", &t)
	var s []int
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
	for i := range s {
		s[i]++
	}
	fmt.Println(s)
}
