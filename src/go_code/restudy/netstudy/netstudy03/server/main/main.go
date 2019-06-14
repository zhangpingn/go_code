package main

import (
	"fmt"
	"net"
	"go_code/restudy/netstudy/netstudy03/common/utils"
	"go_code/restudy/netstudy/netstudy03/common/message"
	"go_code/restudy/netstudy/netstudy03/server/process"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("服务器创建监听失败", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器获取客户端连接失败", err)
			continue
		}
		go dial(conn)
	}
}

func dial(conn net.Conn) {
	var tf *utils.Transfer = &utils.Transfer{
		Conn : conn,
		Buf : make([]byte, 8192),
	}
	mes := tf.ReadPkg()
	switch mes.MsgType {
		case message.LoginMesType:
			var up *process.UserProcess = &process.UserProcess{
				Conn : conn,
			}
			up.Login(mes)
	}
}