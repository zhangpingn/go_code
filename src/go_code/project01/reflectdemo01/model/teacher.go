package model

import (
	_ "fmt"
)

type Teacher struct {
	name string 
	age int
}

func (f Factory) getTeacherFactory(name string, age int) *Teacher {
	return &Teacher{"lisi", 25}
}

// func (tea *Teacher) sayHello1(addr string, weight float64) {
// 	fmt.Printf("学生的姓名是%s,年龄是%d,地址是%s,体重是%.2f", tea.name, tea.age, addr, weight)
// }