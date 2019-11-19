package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	for {
		e, err := r.Read(b)
		fmt.Printf("e = %v err = %v b = %v\n", e, err, b)
		fmt.Printf("b[:e] = %q\n", b[:e])
		if err == io.EOF {
			break
		}
	}
}
