package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Sid uint
	Sname string
	Age uint
}

func main() {
	s := Student{
		Sid : 1001,
		Sname : "张三",
		Age : 20,
	}
	
	str, err := json.Marshal(s)
	if err != nil {
		fmt.Println("格式化对象失败")
		return
	}
	fmt.Println(string(str))

	slice := []string{"zhangsan", "李四", "wangwu"}
	sliceStr, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("格式化对象失败")
		return
	}
	fmt.Println(string(sliceStr))
}