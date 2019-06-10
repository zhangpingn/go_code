package main

import (
	"encoding/xml"
	"fmt"
	_ "os"
	"io/ioutil"
)

type Mapper struct {
	Selectlist []Select `xml:"select"`
	Updatelist []Update	`xml:"update"`
	Insertlist []Insert `xml:"insert"`
	Namespace string `xml:"namespace,attr"`
}
type Select struct {
	Id string `xml:"id,attr"`
	ParameterType string `xml:"parameterType,attr"`
	ResultType string `xml:"resultType,attr"`
	Comment string `xml:",innerxml"`
}
type Update struct {
	Id string `xml:"id,attr"`
	ParameterType interface{} `xml:"parameterType,attr"`
	Comment string `xml:",innerxml"`
}
type Insert struct {
	Id string `xml:"id,attr"`
	ParameterType string `xml:"parameterType,attr"`
	Comment string `xml:",innerxml"`
}
func main() {
	file := "d:/stu.xml"
	comment, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return 
	}
	var mapper Mapper
	xml.Unmarshal([]byte(comment), &mapper)
	//fmt.Println(mapper)
	for k, v := range mapper.Selectlist {
		fmt.Printf("第%d个select的id是%v,入参类型是%v,出参类型是%v,语句是%v\n", k, 
			v.Id, v.ParameterType, v.ResultType, v.Comment)
	}
}