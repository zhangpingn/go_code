package main

import (
	"fmt"
	"net"
	"go_code/restudy/netstudy/netstudy02/common/utils"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"go_code/restudy/netstudy/netstudy02/server/process"
	_ "go_code/restudy/netstudy/netstudy02/common/model"
	"encoding/json"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("创建服务监听失败", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("获取与客户端的连接失败", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var tf = &utils.Transfer{
		Conn:conn,
	}
	for{
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("服务器获取数据失败", err)
			return
		}
		switch mes.MsgType {
			case message.LoginMesType:
				var up *process.UserProcess = &process.UserProcess{
					Conn : conn,
				}
				var loginMes message.LoginMes
				err := json.Unmarshal([]byte(mes.Data), &loginMes)
				if err != nil {
					fmt.Println("反序列化登陆用户数据失败")
					return
				}
				up.Login(&loginMes.User)
			case message.RegisterMesType:
				var registerMes message.RegisterMes
				err := json.Unmarshal([]byte(mes.Data), &registerMes)
				if err != nil {
					fmt.Println("服务器反序列化用户注册信息失败", err)
					return
				}
				var up *process.UserProcess = &process.UserProcess{
					Conn : conn,
				}
				up.Register(&registerMes.User)
			case message.InfoMesType:
				var infoMes message.InfoMes
				fmt.Println(mes.Data)
				err := json.Unmarshal([]byte(mes.Data), &infoMes)
				if err != nil {
					fmt.Println("服务器端反序列化用户信息失败", err)
					return
				}
				var mp *process.MsgProcess = &process.MsgProcess{
					Conn : conn,
				}
				mp.SendMsg(mes, infoMes.ToUserId)
			default:
		}
	}
}