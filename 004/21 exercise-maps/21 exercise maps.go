package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	arr := strings.Fields(s)

	for _, word := range arr {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
}
