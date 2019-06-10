package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{10, 20, 30, 40, 50}
	var slice = make([]int, 2, 2)
	// var slice []int
	fmt.Printf("slice的长度是%d, 容量是%d\n", len(slice), cap(slice))
	fmt.Printf("arr[1]的地址是%p, arr[2]的地址是%p\n", &arr[1], &arr[2])
	fmt.Printf("arr1[1]的地址是%p, arr1[2]的地址是%p\n", &arr1[1], &arr1[2])
	//slice = append(slice, arr[1])
	//fmt.Printf("slice的值为%p\n", slice)
	slice = arr[1:2]
	fmt.Printf("slice的长度是%d, 容量是%d\n", len(slice), cap(slice))
	slice = append(slice, arr[2])
	slice = append(slice, arr1[1])
	slice = append(slice, arr1[1])
	
	// slice2 := arr1[1:3]
	// slice := append(slice1, slice2...)
	//fmt.Printf("类型为%T", slice...)
	//slice = append(slice, arr1[3])
	fmt.Printf("slice的长度是%d, 容量是%d, 地址为%p, 值为%v\n", len(slice), cap(slice), slice, slice)
	slice[1] = 10
	fmt.Printf("slice的值为%v, 数组arr的值为%v\n, slice[0]的地址是%p, slice[1]的地址是%p\n, slice[2]的地址是%p, slice[3]的地址是%p\n", 
		slice, arr, &slice[0], &slice[1], &slice[2], &slice[3])
	//slice[2] = 10
	// var a float64 = 10 + 10
	// fmt.Printf("a=%f", a)
}