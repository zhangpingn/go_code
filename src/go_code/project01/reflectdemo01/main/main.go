package main

import (
	"reflect"
	"go_code/project01/reflectdemo01/model"
	"fmt"
)
func reflectMethod(obj interface{}, methodName string, args... interface{}) interface{}{
	rValue := reflect.ValueOf(obj)
	methodByName := rValue.MethodByName(methodName)
	var argsSlice []reflect.Value
	//便利参数
	for _, v := range args{
		//fmt.Println(v)
		rv := reflect.ValueOf(v)
		argsSlice = append(argsSlice, rv)
	}
	//fmt.Println(argsSlice)
	rv := methodByName.Call(argsSlice)
	var objNew interface{}
	for _, v := range rv{
		objNew = v.Interface()
	}
	return objNew
}

func main() {
	factory := &model.Factory{}
	v := reflectMethod(factory, "GetStudentFactory", "zhansgan", 20)
	fmt.Println(v)
}