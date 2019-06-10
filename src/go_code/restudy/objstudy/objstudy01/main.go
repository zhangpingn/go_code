package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age uint
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func working(s Student) {
	fmt.Println("my name is", s.Name)
}

func main() {
	var p Pupil = Pupil{
		Student{
			Name : "lisi",
			Age : 20,
		},
	}

	p.Name = "zhangsan"

	working(p.Student)
}