package parsexml

import (
	"io/ioutil"
	"encoding/xml"
	_ "go_code/execrise/xmlexec/xmlexec03/model"
	"fmt"
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
	ParameterType string `xml:"parameterType,attr"`
	Comment string `xml:",innerxml"`
}
type Insert struct {
	Id string `xml:"id,attr"`
	ParameterType string `xml:"parameterType,attr"`
	Comment string `xml:",innerxml"`
}
func Parse(file string, sqlId string) interface{}{
	comment, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("文件%v读取失败,err=%v", file, err)
		return nil
	}
	var mapper Mapper
	xml.Unmarshal(comment, &mapper)
	for _, v := range mapper.Selectlist {
		if v.Id == sqlId {
			return v
		}
	}

	for _, v := range mapper.Updatelist {
		if v.Id == sqlId {
			return v
		}
	}

	for _, v := range mapper.Insertlist {
		if v.Id == sqlId {
			return v
		}
	}
	return nil
}