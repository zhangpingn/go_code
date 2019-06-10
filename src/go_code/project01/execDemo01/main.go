package main

import (
	"fmt"
)
func main() {
	var num int = 50
	for i := 1; i <= num; i++ {
		for j := 1; j <= num - i; j++ {
			fmt.Print(" ")
		}
		for j :=1; j <= i * 2 - 1; j++ {
			if(j == 1 || j == i * 2 -1 || i == num) {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
				
		}
		fmt.Println()
	}
	fmt.Scanln()
}