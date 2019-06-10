package process

import (
	"net"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"go_code/restudy/netstudy/netstudy02/common/utils"
	"fmt"
	"encoding/json"
)
type MsgProcess struct {
	Conn net.Conn
}

func (mp *MsgProcess)SendMsg(mes *message.Message, userId string){
	up, ok := OnlineUsers[userId]
	var mesRes message.Message
	mesRes.MsgType = message.InfoMesResType
	var infoMesRes message.InfoMesRes
	var flag bool
	var tf *utils.Transfer = &utils.Transfer{}
	if !ok {
		fmt.Println("消息发送失败")
		infoMesRes.ErrorCode = 500
		infoMesRes.ErrorDesc = "对方不在线"
		flag = true
	}else{
		tf.Conn = up.Conn		
		err := tf.WritePkg(mes)
		if err != nil {
			fmt.Println("服务器发送消息失败", err)
			infoMesRes.ErrorCode = 500
			infoMesRes.ErrorDesc = "信息发送失败"
			flag = true
		}
	}
	
	if !flag {
		infoMesRes.ErrorCode = 200
		infoMesRes.ErrorDesc = "信息发送成功"
	}
	infoMesResMsg, err := json.Marshal(infoMesRes)
	if err != nil {
		fmt.Println("服务器序列化消息失败", err)
		return
	}
	mesRes.Data = string(infoMesResMsg)
	tf.Conn = mp.Conn
	tf.WritePkg(&mesRes)
}