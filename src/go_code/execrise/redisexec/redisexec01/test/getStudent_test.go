package test

import (
	"go_code/execrise/redisexec/redisexec01/utils"
	_ "go_code/execrise/redisexec/redisexec01/model"
	_ "encoding/json"
	"fmt"
	"testing"
)

func TestGetStudent(t *testing.T){
	stu := utils.GetStudent("100001")
	fmt.Println(stu)
}

func TestHGetStudent(t *testing.T){
	utils.HGetStudent("100002")
	//fmt.Println(stu)
}