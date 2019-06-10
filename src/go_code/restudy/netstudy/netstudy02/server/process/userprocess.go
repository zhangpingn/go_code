package process

import (
	"net"
	"go_code/restudy/netstudy/netstudy02/common/model"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"go_code/restudy/netstudy/netstudy02/common/utils"
	"go_code/restudy/netstudy/netstudy02/server/dao"
	"fmt"
	"encoding/json"
)

var OnlineUsers map[string]*UserProcess = make(map[string]*UserProcess, 1)

type UserProcess struct {
	Conn net.Conn
}

func (up *UserProcess) Login(user *model.User) (err error){
	var userDao *dao.UserDao = &dao.UserDao{}
	var mesRes message.Message
	mesRes.MsgType = message.LoginMesResType
	var loginMesRes message.LoginMesRes
	loginUser, err := userDao.GetUserByUserId(user.UserId)
	if err != nil {
		fmt.Println("获取待登录的用户数据失败", err)
		loginMesRes.ErrorCode = 500
		loginMesRes.ErrorDesc = "用户数据不存在"
	}

	tf := &utils.Transfer{}
	if loginUser.Password == user.Password {
		loginMesRes.ErrorCode = 200
		loginMesRes.ErrorDesc = "登录成功"
		n, err := userDao.UpdateStatus(user.UserId, "A")
		if err != nil || n != 1{
			fmt.Println("登陆失败", err)
			return err
		}
		users, _ := userDao.GetAllOnlineUser(user.UserId)
		loginMesRes.Users = users

		var loginMes message.Message
		var onlineMes message.OnlineMes
		loginUser.Password = ""
		loginUser.UserStatus = "A"
		onlineMes.User = *loginUser
		onlineMsg, err := json.Marshal(onlineMes)
		if err != nil {
			fmt.Println("序列化个人上线信息失败")
		}
		loginMes.MsgType = message.OnlineMesType
		loginMes.Data = string(onlineMsg)
		for _, v := range OnlineUsers {
			tf.Conn = v.Conn
			tf.WritePkg(&loginMes)
		}
		OnlineUsers[user.UserId] = up
	}else {
		loginMesRes.ErrorCode = 500
		loginMesRes.ErrorDesc = "用户名或密码错误"
	}
	loginMsg, err := json.Marshal(loginMesRes)
	if err != nil {
		fmt.Println("序列化登陆返回信息失败", err)
		return
	}
	mesRes.Data = string(loginMsg)
	tf.Conn = up.Conn
	tf.WritePkg(&mesRes)
	return 
}

func (up *UserProcess) Register(user *model.User) (err error){
	var userDao *dao.UserDao = &dao.UserDao{}
	_, err = userDao.GetUserByUserId(user.UserId)
	var mesRes *message.Message = &message.Message{}
	mesRes.MsgType = message.RegisterMesResType
	var registerMesRes message.RegisterMesRes
	var flag bool
	if err != nil {
		n, err := userDao.AddUser(user)
		if err == nil && n == 1 {
			registerMesRes.ErrorCode = 200
			registerMesRes.ErrorDesc = "用户注册成功"
			flag = true
		}
	}
	if !flag {
		registerMesRes.ErrorCode = 500
		registerMesRes.ErrorDesc = "用户注册失败"
	}
	regMesRes, err := json.Marshal(registerMesRes)
	if err != nil {
		fmt.Println("用户注册返回数据序列化失败", err)
		return
	}
	mesRes.Data = string(regMesRes)
	var tf *utils.Transfer = &utils.Transfer{
		Conn : up.Conn,
	}
	tf.WritePkg(mesRes)
	return
}