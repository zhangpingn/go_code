package main
import (
	"fmt"
	_"io/ioutil"
)

func main() {

	var map1 map[string] int = make(map[string] int)
	//append(map1, "name", 11)
	map1["age"] = 11
	fmt.Println(map1)
	delete(map1, "age")
	fmt.Println(map1)
	// var s [] int = make([] int, 10, 10)
	// arr := [] int {1,2,3}
	// s[0] = 1
	// fmt.Println(s, &s[0])
	// for _, v := range arr {
	// 	s = append(s, v)
	// }
	
	// fmt.Println(s, &s[0])
	// arr := [3] int {1,2,3}
	// var arr1 [3] int
	// arr1 = arr
	// for k, v := range(arr1) {
	// 	// if k == 1 {
	// 	// 	v = 5
	// 	// }
	// 	fmt.Println(k, v)
	// }
	// for i := 0; i < len(arr1); i++ {
	// 	if i == 1 {
	// 		arr1[1] = 5
	// 	}
	// 	fmt.Println("值为", arr1[i])
	// }
	// for k, v := range(arr1) {
	// 	// if k == 1 {
	// 	// 	v = 5
	// 	// }
	// 	fmt.Println(k, v)
	// }
	// arr := [3][2] int {{1,2},{3,4},{6,5}}
	// for k, v := range arr {
	// 	fmt.Println(k, v)
	// 	for k1, v1 := range v {
	// 		fmt.Println(k1, v1)
	// 	}
	// }
	// content, _ := ioutil.ReadFile("C:\\data\\pluc\\um\\1.txt")
	// fmt.Printf("%s", content)
	// var c = 5
	// fmt.Println(&c)
	// a, b := hello(6, &c)
	// fmt.Printf("%d,%d",a, b)
}
// func hello(a int, b *int) (c,d int) {
// 	// c = a + b 
// 	// d = a - b
// 	fmt.Println(&a, &b)
// 	return
// }