package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "hello~~~````张三"
	for index := 0; index < len(str); index++ {
		fmt.Printf("%c", str[index])
	}
	fmt.Println()
	for _, v := range str {
		fmt.Printf("%c", v)
	}
	label:
	for i:=1; i<10; i++ {
		for j:=0; j<i; j++ {
			if j == 5 {
				continue label
			}
			fmt.Println(i, j)
		}
	}
	str = " hello !!"
	str = strings.Trim(str, " !")
	fmt.Println(str)
	str = "hello.jpg"
	b := strings.HasSuffix(str, ".jpg")
	fmt.Println(b)
	str = strings.ToUpper(str)
	fmt.Println(str)
	fmt.Println(strings.EqualFold("HELLO.JPG",str))
}