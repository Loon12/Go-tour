package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	arr := strings.Fields(s)

	for i := 0; i < len(arr); i++ {
		_, key := m[arr[i]]
		if key == false {
			k := 0
			for j := 0; j < len(arr); j++ {
				if arr[i] == arr[j] {
					k++
				}
			}
			m[arr[i]] = k
		}

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
