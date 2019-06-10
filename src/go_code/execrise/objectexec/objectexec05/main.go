package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Sid string `json:"sid"`
	Sname string `json:"sname"`
}
func getStudent() (err error, stu *Student){
	var stuMsg = "{\"sid\":\"10001\",\"sname\":\"zhangsan\"}"
	stu = &Student{}
	err = json.Unmarshal([]byte(stuMsg), stu)
	return
}

func main() {
	err, stu := getStudent()
	if err != nil {
		fmt.Println("序列化有误", err)
		return
	}
	fmt.Println(*stu)

}