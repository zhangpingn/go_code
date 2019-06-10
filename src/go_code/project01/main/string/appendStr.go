package main

import (
	"strconv"
	"fmt"
)
func main() {

	//var str string = "hello"
	var numbers [] string
	numbers = append(numbers, strconv.Itoa(4))
	for i :=0; i < len(numbers); i++ {
		//numbers = append(numbers, str, strconv.Itoa(i))
		fmt.Println(numbers[i])
	}
	var numbers1 [1][1] string
	//numbers1[0] = append(numbers1[0], str)
	fmt.Println(len(numbers1))
	// for i:= 0; i < len(numbers1); i++ {
	// 	for j :=0; j < len(numbers1[i]); j++ {
	// 		fmt.Println(numbers1[i][j])
	// 	}
	// }
}