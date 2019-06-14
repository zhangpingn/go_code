package process

import (
	"net"
	"go_code/restudy/netstudy/netstudy03/common/message"
	"encoding/json"
	"fmt"
	"go_code/restudy/netstudy/netstudy03/server/dao"
	"go_code/restudy/netstudy/netstudy03/common/utils"
	"go_code/restudy/netstudy/netstudy03/common/model"
)
var HeroesMap map[string]model.Heroes = make(map[string]model.Heroes, 1)
type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) Login(mes *message.Message) {
	var loginMes message.LoginMes
	var loginMesRes message.LoginMesRes
	err := json.Unmarshal([]byte(mes.Content), &loginMes)
	if err != nil {
		fmt.Println("服务器反序列化客户登录数据失败", err)
		return
	}
	fmt.Println(loginMes.User)
	mes.MsgType = message.LoginMesResType
	var ud dao.UserDao
	var user = ud.Login(loginMes.User.UserId)
	fmt.Println(user)
	if user.Password == loginMes.User.Password {
		var hd dao.HeroDao
		var hero = hd.GetHeroes(user.UserName)
		hero.BloodChan = make(chan uint, 10)
		loginMesRes.ErrorCode = 200
		loginMesRes.ErrorDesc = "登录成功"
		loginMesRes.HeroName = (*hero).HeroName
		loginMesRes.Blood = (*hero).Blood
		loginMesRes.Power = (*hero).Power
		HeroesMap[loginMesRes.HeroName] = *hero
	}else {
		loginMesRes.ErrorCode = 500
		loginMesRes.ErrorDesc = "登录失败，用户名或密码不存在"
	}
	var tf *utils.Transfer = &utils.Transfer{
		Conn : this.Conn,
	}
	fmt.Println(loginMesRes)
	loginMesResMsg, err := json.Marshal(loginMesRes)
	if err != nil {
		fmt.Println("服务器序列化客户登录返回数据失败", err)
		return
	}
	mes.Content = string(loginMesResMsg)
	tf.WritePkg(mes)
}