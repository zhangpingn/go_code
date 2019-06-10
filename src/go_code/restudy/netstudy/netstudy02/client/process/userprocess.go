package process

import (
	"net"
	"fmt"
	"go_code/restudy/netstudy/netstudy02/common/model"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"go_code/restudy/netstudy/netstudy02/common/utils"
	"encoding/json"
	"os"
)

var UserId string
var Conn net.Conn
var onlineUsers map[string]*model.User = make(map[string]*model.User)
type UserProcess struct {
}

func (up *UserProcess) Login(){
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("获取与服务器的连接失败", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.MsgType = message.LoginMesType
	var loginMes message.LoginMes
	var user model.User
	fmt.Println("请输入userId")
	fmt.Scanln(&user.UserId)
	fmt.Println("请输入用户密码")
	fmt.Scanln(&user.Password)
	loginMes.User = user
	msg, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化登录消息失败", err)
		return
	}
	mes.Data = string(msg)
	var tf = &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("用户登录数据发送失败")
		return
	}
	mesRes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("获取用户登录返回数据失败")
		return
	}

	if mesRes.MsgType == message.LoginMesResType {
		var loginMesRes message.LoginMesRes
		err = json.Unmarshal([]byte(mesRes.Data), &loginMesRes)
		if err != nil {
			fmt.Println("序列化登录返回数据失败", err)
			return
		}
		if loginMesRes.ErrorCode == 200 {
			UserId = user.UserId
			Conn = conn
			fmt.Println("用户登录成功")
			fmt.Println("在线用户列表如下")
			for k, v := range loginMesRes.Users {
				fmt.Print(v.UserName, " ")
				onlineUsers[v.UserId] = &v
				if k % 4 == 0 {
					fmt.Println()
				}
			}
			fmt.Println()
			go getMsg(conn)
			menu()
		}else if loginMesRes.ErrorCode == 500 {
			fmt.Println(loginMesRes.ErrorDesc)
			return
		}
	}
}

func (up *UserProcess) Register(){
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("获取服务器连接失败")
		return
	}
	defer conn.Close()
	var mes *message.Message = &message.Message{}
	mes.MsgType = message.RegisterMesType
	var registerMes message.RegisterMes
	var user model.User = model.User{
		UserStatus : "C",
	}
	fmt.Println("请输入您的用户Id")
	fmt.Scanln(&user.UserId)
	fmt.Println("请输入您的用户名")
	fmt.Scanln(&user.UserName)
	fmt.Println("请输入您的密码")
	fmt.Scanln(&user.Password)
	registerMes.User = user
	regMes, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化用户注册信息失败", err)
		return
	}
	mes.Data = string(regMes)
	var tf *utils.Transfer = &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(mes)
	if err != nil {
		fmt.Println("发送用户注册信息失败", err)
		return
	}
	mes, err = tf.ReadPkg()
	if mes.MsgType == message.RegisterMesResType{
		var registerMesRes message.RegisterMesRes
		err := json.Unmarshal([]byte(mes.Data), &registerMesRes)
		if err != nil {
			fmt.Println("反序列化用户注册返回数据失败", err)
			return
		}
		if registerMesRes.ErrorCode == 200 {
			fmt.Println("用户注册成功")
		}else if registerMesRes.ErrorCode == 500 {
			fmt.Println("用户注册失败", registerMesRes.ErrorDesc)
		}
	}
}

func menu() {
	var key string
	for {
		fmt.Println("-----------1、发送消息-------------")
		fmt.Println("-----------2、显示用户列表---------")
		fmt.Println("-----------3、退出系统-------------")
		fmt.Println("请选择")
		fmt.Scanln(&key)
		switch key {
			case "1":
				var mp *MsgProcess = &MsgProcess{
					Conn : Conn,
				}
				mp.SendMsg()
			case "2":
				fmt.Println("显示用户列表")
			case "3":
				os.Exit(0)
		}
	}
}

func getMsg(conn net.Conn){
	var tf *utils.Transfer = &utils.Transfer{
		Conn : conn,
	}
	for{
		mesRes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("获取服务器返回数据失败")
		}
		fmt.Println(mesRes.MsgType, mesRes.Data)
		switch mesRes.MsgType {
			case message.OnlineMesType:
				var onlineMes message.OnlineMes
				fmt.Println(mesRes.Data)
				err := json.Unmarshal([]byte(mesRes.Data), &onlineMes)
				if err != nil {
					fmt.Println("反序列化服务器返回的数据失败", err)
					break
				}
				var user = onlineMes.User
				fmt.Println(user.UserName, "上线了")
				onlineUsers[user.UserId] = &user
			case message.InfoMesType:
				var infoMes message.InfoMes
				err := json.Unmarshal([]byte(mesRes.Data), &infoMes)
				if err != nil {
					fmt.Println("反序列化用户发送的信息失败", err)
					break
				}
				fmt.Println(infoMes.FromUserId, "对您发了一个消息", infoMes.Info)
			case message.InfoMesResType:
				var infoMesRes message.InfoMesRes
				err = json.Unmarshal([]byte(mesRes.Data), &infoMesRes)
				if err != nil || infoMesRes.ErrorCode == 500 {
					fmt.Println("发送消息失败", infoMesRes.ErrorDesc)
					break
				}
				fmt.Println("发送消息成功")
		}
	}
}