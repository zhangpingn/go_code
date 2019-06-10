package process

import (
	"fmt"
	"go_code/restudy/netstudy/netstudy03/common/model"
	"go_code/restudy/netstudy/netstudy03/common/message"
	"encoding/json"
)

type UserProcess struct {

}

func (this *UserProcess) Login() {
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
}