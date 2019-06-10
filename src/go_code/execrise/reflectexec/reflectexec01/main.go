package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age uint
	Score float64
}

func test(a interface{}){
	val := reflect.ValueOf(a)
	val.Elem().Field(0).SetString("lisi")
	val.Elem().Field(1).SetUint(20)
	val.Elem().Field(2).SetFloat(98.2)
	// val.Field(0).SetString("lisi")
	// val.Field(1).SetUint(20)
	// val.Field(2).SetFloat(98.2)
}

func main() {
	var stu Student = Student{
		Name:"zhangsan",
		Age:30,
		Score:28.0,
	}
	fmt.Println(stu)
	test(&stu)
	fmt.Println(stu)
}