package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 20
	str := strconv.Itoa(num1)
	fmt.Printf("str的类型是%T,值是%s\n", str, str)
	var ptr *int = &num1
	*ptr = 30
	num1 = 40
	fmt.Println(num1, *ptr)
}