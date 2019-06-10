package main

import (
	"fmt"
	"time"
	"sync"
)

var num []int
var lock sync.Mutex

func test(i int) {
	//lock.Lock()
	var res int = 1
	for j:=2; j<=i; j++ {
		res *= j
	}
	//num[i] = res
	num = append(num, res)
	//lock.Unlock()
}

func main(){
	num = make([]int, 0)
	for i:=0; i<60; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 3)
	for k, v := range num {
		fmt.Printf("第%v个值是%v\n", k+1, v)
	}
}