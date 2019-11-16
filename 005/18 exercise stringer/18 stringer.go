package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s *struct{} = nil
	fmt.Println(unsafe.Sizeof(s))
	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i))
	var a []int = nil
	fmt.Println(unsafe.Sizeof(a))
	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m))
	var f func() = nil
	fmt.Println(unsafe.Sizeof(f))
	var b []bool = nil
	fmt.Println(unsafe.Sizeof(b))
}
