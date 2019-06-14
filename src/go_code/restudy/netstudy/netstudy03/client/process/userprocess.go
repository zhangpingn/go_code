package process

import (
	"fmt"
	"go_code/restudy/netstudy/netstudy03/common/model"
	"go_code/restudy/netstudy/netstudy03/common/message"
	"encoding/json"
	"go_code/restudy/netstudy/netstudy03/common/utils"
	"net"
)

var Hero model.Heroes
type UserProcess struct {

}

func (this *UserProcess) Login() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("客户端连接服务器失败", err)
		return
	}
	var mes message.Message
	var loginMes message.LoginMes
	mes.MsgType = message.LoginMesType
	var user model.User
	fmt.Println("请输入UserId")
	fmt.Scanln(&user.UserId)
	fmt.Println("请输入密码")
	fmt.Scanln(&user.Password)
	loginMes.User = user
	loginMesMsg, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化用户登录信息失败")
		return
	}
	mes.Content = string(loginMesMsg)
	var tf *utils.Transfer = &utils.Transfer{
		Conn : conn,
		Buf : make([]byte, 8192),
	}
	tf.WritePkg(&mes)
	var mesRes = tf.ReadPkg()
	fmt.Println(mesRes)
	if mesRes.MsgType == message.LoginMesResType {
		var loginMesRes message.LoginMesRes
		err = json.Unmarshal([]byte(mesRes.Content), &loginMesRes)
		if err != nil {
			fmt.Println("反序列化登录返回的数据失败", err)
			return
		}
		if loginMesRes.ErrorCode == 200 {
			fmt.Println("登录成功")
			fmt.Println("您的信息是", loginMesRes.HeroName, loginMesRes.Blood, loginMesRes.Power)
			Hero.HeroName = loginMesRes.HeroName
			Hero.Blood = loginMesRes.Blood
			Hero.Power = loginMesRes.Power
		}else if loginMesRes.ErrorCode == 500 {
			fmt.Println("登录失败", loginMesRes.ErrorDesc)
		}
	}
}

func Menu() {
	var key string
	for {
		fmt.Println("-------------请选择您的操作----------")
		fmt.Println("---------------1、攻击--------------")
		fmt.Println("-------------2、显示信息------------")
		fmt.Scanln(&key)
		switch key {
			case "1":
				var attackMes message.AttackMes
				fmt.Println("请输入您要攻击目标的名字")
				fmt.Scanln(&attackMes.ToHeroName)

			case "2":
		}
	}
}