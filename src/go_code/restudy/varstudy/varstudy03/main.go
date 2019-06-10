package main

import (
	"fmt"
	"time"
	_ "errors"
)
func test(){
	defer func (){
		err := recover()
		fmt.Println(err)
	}()
	panic("hello")
}
func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Year(), int(now.Month()), now.Day(), now.Hour(), 
	   now.Minute(), now.Second())
	test()
	fmt.Println("hello11")

	var arr [5]int8 = [...]int8{10, 5, 8, 10, 2}
	fmt.Printf("%p, %p, %p, %p", &arr, &arr[0], &arr[1], &arr[2])
	var slice = arr[0:3]
	fmt.Println("\n", cap(slice), len(slice))

	slice = make([]int8, 5, 10)
	fmt.Printf("\n%p", slice)
	slice = append(slice, 8)
	fmt.Printf("\n%p", slice)
	for i := 0; i<5; i++{
		slice = append(slice, 8)
		fmt.Printf("\n%p", slice)
	}
	var slice1 = []int{10, 20, 30, 40, 50}
	var slice2 = make([]int, 10)
	copy(slice2, slice1)
	slice1[1] = 100
	fmt.Println(slice2)
}