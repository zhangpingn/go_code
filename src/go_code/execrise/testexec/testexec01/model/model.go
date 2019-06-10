package model

import (
	"fmt"
)

func GetSum() bool{
	var sum int
	for i:=0; i<10; i++ {
		sum += i
	}
	if sum == 45 {
		fmt.Println("结果是", sum)
		return true
	}
	return false
}

func GetSub(n1, n2, res int) bool {
	if res == (n1 - n2) {
		return true
	}else {
		return false
	}
}