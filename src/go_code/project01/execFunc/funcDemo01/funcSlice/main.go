package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Age uint `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	var slice = []int{2, 3, 4, 5}
	fmt.Printf("修改之前的slice地址%p\n", slice)

	slice = append(slice, 6)
	fmt.Printf("修改之后的slice地址%p\n", slice)

	var map1 map[string]string = map[string]string{
		"name":"zhangsan",
		"age":"16",
	}
	fmt.Printf("map1的地址%p，值为%v\n", map1, map1)

	stu := Student{"tom", 20, []string{"play"}}
	resStu,_ := json.Marshal(stu)
	fmt.Printf("序列号之后的resStu的值为%v\n", string(resStu))
	
	var str string = "hello" + "123" +
	"hello" + "123"
	fmt.Println(str)

	var stuPtr *Student = &Student{"jerry", 20, []string{"play"}}
	fmt.Printf("stuPtr的类型为%T", stuPtr)
}