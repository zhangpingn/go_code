package utils

import (
	"net"
	"go_code/restudy/netstudy/netstudy03/common/message"
	"encoding/json"
	"fmt"
	"encoding/binary"
)

type Transfer struct {
	Conn net.Conn
	Buf []byte
}

func (this *Transfer) WritePkg(mes *message.Message) {
	mesMsg, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化待发送的消息失败")
		return
	}
	var length uint32 = uint32(len(mesMsg))
	var buf []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(buf[:4], length)
	this.Conn.Write(buf)
	this.Conn.Write(mesMsg)
}

func (this *Transfer) ReadPkg() (mesRes *message.Message){
	n, err := this.Conn.Read(this.Buf[:4])
	if err != nil || n != 4{
		fmt.Println("接收信息失败")
		return
	}
	var mesMsgLength uint32 = binary.BigEndian.Uint32(this.Buf[:4])
	n, err = this.Conn.Read(this.Buf[:mesMsgLength])
	if err !=nil || uint32(n) != mesMsgLength {
		fmt.Println("接收信息内容失败")
		return
	}
	mesRes = &message.Message{}
	err = json.Unmarshal(this.Buf[:mesMsgLength], &mesRes)
	if err != nil {
		fmt.Println("反序列化信息失败")
		return
	}
	return mesRes
}