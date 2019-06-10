package process

import (
	"fmt"
	"net"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"go_code/restudy/netstudy/netstudy02/common/utils"
	"encoding/json"
)

type MsgProcess struct {
	Conn net.Conn
}

func (mp *MsgProcess) SendMsg(){
	var mes *message.Message = &message.Message{}
	mes.MsgType = message.InfoMesType
	var infoMes message.InfoMes
	fmt.Println("请输入要发送消息人的UserId")
	fmt.Scanln(&infoMes.ToUserId)
	fmt.Println("请输入要发送的消息内容")
	fmt.Scanln(&infoMes.Info)
	infoMes.FromUserId = UserId
	msg, err := json.Marshal(infoMes)
	if err != nil {
		fmt.Println("序列化待发送的消息失败", err)
		return
	}
	mes.Data = string(msg)
	var tf *utils.Transfer = &utils.Transfer{
		Conn : mp.Conn,
	}
	err = tf.WritePkg(mes)
	if err != nil {
		fmt.Println("发送消息失败")
		return
	}
}