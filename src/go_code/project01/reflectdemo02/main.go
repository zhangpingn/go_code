package main

import(
	"fmt"
	"reflect"
)

func main() {
	var num = 10
	rval := reflect.ValueOf(num)
	fmt.Printf("rval的类型为%T,类别为%v,值为%v", rval, rval.Kind(), rval)
	rval.SetInt(20)
	fmt.Println(num)
}