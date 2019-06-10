package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for k, v := range args {
		fmt.Printf("第%v个参数是%v\n", k, v)
	}
}