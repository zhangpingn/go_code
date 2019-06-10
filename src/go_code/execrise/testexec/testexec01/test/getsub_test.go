package test

import (
	"go_code/execrise/testexec/testexec01/model"
	"testing"
)

func TestGetSub(t *testing.T){
	flag := model.GetSub(10, 7, 3)
	if flag {
		t.Logf("答案正确")
	}else {
		t.Fatalf("答案错误")
	}

}