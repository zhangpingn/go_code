package main

import (
	"fmt"
	"net"
	"os"
	"go_code/restudy/netstudy/netstudy01/client/utils"
)

func main() {
	var key string
	for {
		fmt.Println("--------------------请选择您的输入----------------")
		fmt.Println("-----------------------1.登 录-------------------")
		fmt.Println("-----------------------2.注 册-------------------")
		fmt.Println("-----------------------3.退 出-------------------")
		fmt.Println("请输入您的选择")
		fmt.Scanln(&key)

		switch key {
			case "1" : 
				fmt.Println("登录")
				conn, err := net.Dial("tcp", "localhost:8889")
				if err != nil {
					fmt.Println("获取连接失败", err)
					return
				}
				defer conn.Close()
				utils.Login(conn)
			case "2" :
				fmt.Println("注册用户")
			case "3" :
				os.Exit(0)
			default : 
			 	fmt.Println("输入有误")
 		}
	}
}

func hello(){
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("获取连接失败", err)
		return
	}
	defer conn.Close()

	

	_, err = conn.Write([]byte("hello, world"))
	if err != nil {
		fmt.Println("内容发送失败", err)
	}
}