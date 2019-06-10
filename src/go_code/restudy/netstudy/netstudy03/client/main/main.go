package main 

import (
	"fmt"
)

func main() {
	var key string
	for {
		fmt.Println("---------请输入您的选择--------")
		fmt.Println("-----------1、登录------------")
		fmt.Println("-----------2、注册------------")
		fmt.Println("-----------3、退出------------")
		fmt.Println("请输入")
		fmt.Scanln(&key)
		switch key {
		case "1":
		case "2":
		case "3":
		default:
		}
	}
}