package test

import (
	"go_code/execrise/redisexec/redisexec01/utils"
	"go_code/execrise/redisexec/redisexec01/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetStudent(t *testing.T){
	var sid = "100001"
	stu := model.Student{
		Sid : sid,
		Sname : "zhangsan",
		Age : 20,
		Score : 98.0,
	}
	str, err := json.Marshal(&stu) 
	if err != nil {
		fmt.Println("序列化stu失败")
		return
	}
	stuStr := string(str)
	err = utils.SetStudent(sid, &stuStr)
	if err != nil {
		fmt.Println("设置数据失败")
	}
}

func TestHSetStudent(t *testing.T){
	var sid = "100002"
	var student *model.Student = &model.Student{
		Sid : sid,
		Sname : "lisi",
		Age : 20,
		Score : 99.0,
	}
	err := utils.HSetStudent(sid, student)
	if err != nil {
		t.Fatalf("HSet失败, err=%v", err)
		return
	}
}