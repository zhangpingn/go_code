package model

import (
	"fmt"
)
type student struct{
	name string
	age uint
	score float64
}

func NewStudent(name string, age uint, score float64) *student{
	return &student{
		name : name,
		age : age,
		score : score,
	}
}

func (stu *student) GetName() string{
	return stu.name
}

func (stu *student) SetName(name string){
	stu.name = name
}

func (stu *student) GetAge() uint{
	return stu.age
}

func (stu *student) SetAge(age uint){
	stu.age = age
}

func (stu *student) GetScore() float64{
	return stu.score
}

func (stu *student) SetScore(score float64){
	stu.score = score
}

func (stu *student) String() string{
	return fmt.Sprintf("学生的name是%v, age是%v, score是%v", stu.name, stu.age, stu.score)
}
