package main

import (
	"fmt"
	"go_code/execrise/objectexec/objectexec03/model"
)

func main() {

	var stu = model.NewStudent("zhangsan", 20, 43.2)
	fmt.Println("stu的name是", stu.GetName())
	fmt.Println("stu的age是", stu.GetAge())
	fmt.Println("stu的score是", stu.GetScore())
	stu.SetName("李四")
	stu.SetAge(30)
	stu.SetScore(11.2)
	fmt.Println("stu的name是", stu.GetName())
	fmt.Println("stu的age是", stu.GetAge())
	fmt.Println("stu的score是", stu.GetScore())
}