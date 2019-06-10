package utils

import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/restudy/netstudy/netstudy01/common/message"
)

var users map[string]net.Conn = make(map[string]net.Conn, 1)

//处理客户端发送的请求数据
func Process(conn net.Conn){
	defer conn.Close()
	var buf []byte = make([]byte, 1096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取数据失败", err)
	}
	fmt.Printf("读取的数据长度为%d, 内容为%v", n, string(buf[:n]))
	var mes message.Message
	err = json.Unmarshal(buf[:n], &mes)
	if err != nil {
		fmt.Println("获取消息失败", err)
		return
	}
	switch mes.MesType {
		case message.LoginType:
			var loginMes message.LoginMessage
			err = json.Unmarshal([]byte(mes.Data), &loginMes)
			if err != nil {
				fmt.Println("反序列化消息失败", err)
			}
			loginMesRes := Login(&loginMes, conn)
			buf, err = json.Marshal(loginMesRes)
			if err != nil {
				fmt.Println("序列化登录返回消息失败", err)
			}
			conn.Write(buf)
	}
}

func Login(loginMes *message.LoginMessage, conn net.Conn) (mes message.Message){
	mes.MesType = message.LoginResType
	var loginResMes message.LoginMessageRes
	if (loginMes.UserId == "zhangsan" && loginMes.Pwd == "123456") || (loginMes.UserId == "lisi" && loginMes.Pwd == "123456") {
		loginResMes.Code = 200
		loginResMes.Desc = "登录成功"
		
		users[loginMes.UserId] = conn
		sendOnlineToOthers(loginMes.UserId)
	}else {
		loginResMes.Code = 500
		loginResMes.Desc = "用户名或密码错误"
	}
	buf, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化登录返回消息失败", err)
	}
	mes.Data = string(buf)
	return
}

func sendOnlineToOthers(userId string) {
	for k, v := range users {
		if k == userId {
			continue
		}
		var mes message.Message
		var onlineMessage message.OnlineMessage
		mes.MesType = message.OnlineMessageType
		onlineMessage.UserId = userId
		msg, err := json.Marshal(onlineMessage)
		fmt.Println("序列化用户上线子信息")
		if err != nil {
			fmt.Println("序列化用户上线子信息失败", err)
			return
		}
		mes.Data = string(msg)
		msg, err = json.Marshal(mes)
		fmt.Println("序列化用户上线主信息")
		if err != nil {
			fmt.Println("序列化用户上线主信息失败", err)
			return
		}
		fmt.Println(string(msg))
		v.Write(msg)

	}
	return
}