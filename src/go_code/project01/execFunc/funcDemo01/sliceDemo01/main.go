package main

import (
	"fmt"
)

func main() {
	// var intArr = [...]int8 {10, 20, 30, 21}
	// fmt.Printf("%p\n", &intArr[1])
	// slice := intArr[1:3]
	// fmt.Printf("%p\n", &slice[0])
	// //fmt.Printf("%p\n", slice)
	// fmt.Println(cap(slice))
	// slice = append(slice, 22)
	// slice = append(slice, 23)
	// fmt.Println(cap(slice))
	// // var ptr *int8 
	// // ptr = 0xc00004e059
	// // var slice1 []int = make([]int, 4, 5)
	// // fmt.Println(slice1)
	// // fmt.Println(len(slice1))
	// // fmt.Println(cap(slice1))
	// // slice1 = append(slice1, 10)
	// // fmt.Println(len(slice1))
	// // fmt.Println(cap(slice1))
	// var slice1 []int
	// fmt.Printf("%p", slice1)
	// slice1 = append(slice1, 10)
	// fmt.Println(slice1)
	// funcC := funcB(10, 20)
	// sum := funcC(30, 40)
	// fmt.Println(sum)
	// var str = "hello"
	// slice2 := str[0:]
	// fmt.Printf("%T,%v,%p\n", slice2, len(slice2), &str)
	// var slice3 []byte = []byte(str)
	// slice3[0] = 'e'
	// fmt.Printf("%T,%v,%p\n", slice3, len(slice3), &str)
	// var slice4 []uint = make([]uint, 0)
	// fmt.Printf("slice4==%p\n", slice4)
	// fbn := fbn(10, slice4)
	// fmt.Println("slice4==", slice4)
	// fmt.Printf("%v", fbn)

	var strSlice = []string {"a", "b", "c"}
	for k, v := range strSlice {
		fmt.Println(v, k)
	}
}

var funcA = func (a, b int) (sum int){
	sum = a + b
	return
}

func funcB(a, b int) func (int, int) int {
	return func(c int, d int) int{
		fmt.Println(c, d)
		return a + b
	}
}

func fbn(n int, slice []uint) uint {
	fmt.Printf("slice4==%p\n", slice)
	if n == 1 || n == 2 {
		slice = append(slice, 1)
		return 1
	}else {
		slice = append(slice, fbn(n-1, slice) + fbn(n-2, slice))
		return fbn(n-1, slice) + fbn(n-2, slice)
	}
} 