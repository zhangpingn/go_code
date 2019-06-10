package main

import (
	"fmt"
	"time"
)

var (
	nummap map[int]int
)

func test(j int){
	var res int
	for i:=1; i<=j; i++ {
		res += i
	}
	nummap[j] = res
	time.Sleep(time.Millisecond * 20)
}

func test01(i int){
	fmt.Println("hello, world", i)
}
func main() {
	nummap = make(map[int]int)
	for i:=1; i<20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	for k, v := range nummap {
		fmt.Printf("第%v个数是%v\n", k, v)
	}
	//fmt.Println("hello, golang")
}