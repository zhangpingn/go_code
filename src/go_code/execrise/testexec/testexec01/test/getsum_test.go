package test

import (
	"go_code/execrise/testexec/testexec01/model"
	"testing"
)

func TestGetSum(t *testing.T){
	flag := model.GetSum()
	if flag {
		t.Logf("答案正确")
	}else {
		t.Fatalf("答案错误")
	}

}