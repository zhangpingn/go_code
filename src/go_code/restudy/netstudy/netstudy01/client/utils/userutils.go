package utils

import (
	"fmt"
	"go_code/restudy/netstudy/netstudy01/common/message"
	"encoding/json"
	"net"
	"os"
)
func Login(conn net.Conn){

	var mes message.Message
	var loginMes message.LoginMessage

	mes.MesType = message.LoginType
	fmt.Println("请输入用户名")
	fmt.Scanln(&loginMes.UserId)
	fmt.Println("请输入密码")
	fmt.Scanln(&loginMes.Pwd)
	buf, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化对象失败", err)
	}
	mes.Data = string(buf)
	buf, err = json.Marshal(mes)
	conn.Write(buf)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("获取服务器返回登录信息失败")
	}
	json.Unmarshal(buf[:n], &mes)
	if mes.MesType == message.LoginResType {
		var loginResMes message.LoginMessageRes
		json.Unmarshal([]byte(mes.Data), &loginResMes)
		if loginResMes.Code == 200 {
			fmt.Println("登录成功")
			go getAllMes(conn)
			menu()
		}else{
			fmt.Println("登录失败")
		}
	}
}

func getAllMes(conn net.Conn){
	fmt.Println("进入获取相应的上线用户")
	for {
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		var mes message.Message
		fmt.Println(string(buf[:n]))
		err := json.Unmarshal(buf, &mes)
		if err != nil{
			fmt.Println("反序列化监听的信息失败")
			return
		}
		switch mes.MesType {
			case message.OnlineMessageType : 
				var onlineMessage message.OnlineMessage
				err = json.Unmarshal([]byte(mes.Data), onlineMessage)
				if err !=nil {
					fmt.Println("获取新用户上线信息失败")
					break
				}
				fmt.Println("用户", onlineMessage.UserId, "上线了")
		}
	}
}

func menu(){
	var key string
	for{
		fmt.Println("---------------发送消息-----------")
		fmt.Println("-----------1、显示用户列表---------")
		fmt.Println("-----------2、退出系统-------------")
		fmt.Println("请输入您的选择")
		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("显示用户列表")
			case "2":
				fmt.Println("退出系统")
				os.Exit(0)
		}
	}
}