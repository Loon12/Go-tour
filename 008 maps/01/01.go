package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Так говорила в июле 1805 года известная Анна Павловна Шерер, фрейлина и приближенная императрицы Марии Феодоровны, встречая важного и чиновного князя Василия, первого приехавшего на ее вечер. Анна Павловна кашляла несколько дней, у нее был грипп, как она говорила (грипп был тогда новое сло0во, употреблявшееся только редкими)."
	g := strings.Split(s, " ")
	m := make(map[string]int)
	for i := range g {
		if _, word := m[g[i]]; word == false {
			for j := range g {
				if g[i] == g[j] {
					m[g[i]]++
				}
			}
		}
	}
	fmt.Println(m)

}
