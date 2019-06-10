package model

import (
	_ "fmt"
)

type Student struct {
	name string
	age int
}

type Factory struct {

}

func (f Factory) GetStudentFactory(name string, age int) *Student{
	return &Student{name, age}
}

// func (stu *Student) sayHello(addr string, tail uint, weight float64) {
// 	fmt.Printf("学生的姓名是%s,年龄是%d,地址是%s,身高是%d,体重是%.2f", stu.name, stu.age, addr, tail, weight)
// }