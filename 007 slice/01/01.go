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
		//var t1 int
		//fmt.Printf("Введите элемент массива №%v: ", i)
		//fmt.Scanf("%f \n" & t1)
		s = append(s, i)
	}
	for i := range s {
		s[i]++
	}
	fmt.Println(s)
}
