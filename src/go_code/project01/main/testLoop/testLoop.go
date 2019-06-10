package main

import "fmt"

var number = [] int {1, 2, 3, 4, 5, 6}

func main() {
	// i := 0
	// for i < 6 {
	// 	fmt.Println(number[i])
	// 	i++
	// }

	// for i := 0; i < 6; i++ {
	// 	fmt.Println(number[i])
	// }

	for i,x := range number {
		fmt.Println("第", i, "个数为", x)
	}
}