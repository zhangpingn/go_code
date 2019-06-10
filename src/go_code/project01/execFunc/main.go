package main

import (
	"fmt"
)

func getSum(num1, num2, num3 int) (int, int) {
	return num1 + num2, num1 - num2
}

func getSum1(args... int) (sum int) {
	for _, v := range args{
		sum += v
	}
	return
}

func main() {
	a := getSum
	//fmt.Println(&getSum, &a)
	_, _ = a(10, 5, 8)
	fmt.Println("hello")
	sum := getSum1()
	fmt.Println(sum)
}