package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

type Student struct {
	Name string `xml:"Name"`
	Age int `xml:"Age"`
}

func main() {
	file := "d:/stu.xml"
	commont, ok := ioutil.ReadFile(file)
	if ok != nil {
		fmt.Println("文件读取失败")
		return 
	}
	var stu Student
	xml.Unmarshal(commont, &stu)
	fmt.Println(string(commont))
	fmt.Println(stu)
}