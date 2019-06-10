package main

import (
	"fmt"
)

func main() {
	//var a = 'A'
	var array [26]byte
	//fmt.Printf("a的类型是%T", a)
	for i:=0; i<26; i++ {
		array[i] = 'A' + byte(i)
	}

	for i:=0; i<len(array); i++ {
		fmt.Printf("第%d个元素是%c\n", i, array[i])
	}
	
	var n1 = 19.2
	var n2 = n1 + 10
	fmt.Printf("n2的值为%f", n2)
}