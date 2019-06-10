package main

import (
	"fmt"
)

func main() {
	var str = "大家好"
	b := []rune(str)
	fmt.Printf("str的地址是%p", &str)

	b[0] = '你'
	fmt.Printf("str的地址是%p, 值为%v", &str, str)
}