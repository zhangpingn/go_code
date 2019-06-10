package main

import (
	"encoding/json"
	"fmt"
)
type Student struct{
	Sid uint
	Sname string
	Age uint
}

func test() string{
	return "{\"Sid\":1002,\"Sname\":\"lisi\",\"Age\":20}"
}

func main() {
	str := "{\"Sid\":1001,\"Sname\":\"张三\",\"Age\":20}";
	var stu Student
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("类型转换错误")
		return
	}
	fmt.Println(stu)

	str = "[\"zhangsan\",\"李四\",\"wangwu\"]"
	var slice []string
	json.Unmarshal([]byte(str), &slice)
	fmt.Println(slice)

	str = test()
	fmt.Println(str)
	err = json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("类型转换错误")
		return
	}
	fmt.Println(stu)
}