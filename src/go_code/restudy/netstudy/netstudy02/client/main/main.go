package main

import (
	"fmt"
	"go_code/restudy/netstudy/netstudy02/client/process"
)

func main() {
	var key string
	for {
		fmt.Println("-------------欢迎来到多人聊天系统-----------")
		fmt.Println("-----------------1、登 录------------------")
		fmt.Println("-----------------2、注 册------------------")
		fmt.Println("-----------------3、退 出------------------")
		fmt.Println("请输入您的选择")
		fmt.Scanln(&key)
		switch key {
			case "1":
				var up *process.UserProcess = &process.UserProcess{}
				up.Login()
			case "2":
				var up *process.UserProcess = &process.UserProcess{}
				up.Register()
			case "3":
			default:
				fmt.Println("您的输入有误")
		}
	}
}