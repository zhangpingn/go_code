package utils

import (
	"net"
	"go_code/restudy/netstudy/netstudy02/common/message"
	"encoding/json"
	"fmt"
	"encoding/binary"
)

type Transfer struct {
	Conn net.Conn
	Buf [8192]byte
}

func (tf *Transfer)WritePkg(mes *message.Message)(err error){
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化待发送的消息失败", err)
		return
	}
	var buf [4]byte
	var length uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], length)
	n, err := tf.Conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("发送数据长度失败", err)
		return
	}
	n, err = tf.Conn.Write(data)
	if uint32(n) != length || err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	fmt.Println("发送数据成功", string(data))
	return
}

func (tf *Transfer) ReadPkg() (mesRes *message.Message, err error){
	n, err := tf.Conn.Read(tf.Buf[:4])
	if err != nil || n != 4{
		fmt.Println("读取返回数据长度失败", err)
		return
	}
	length := binary.BigEndian.Uint32(tf.Buf[:4])
	n, err = tf.Conn.Read(tf.Buf[:length])
	if n != int(length) || err != nil {
		fmt.Println("读取返回数据失败", err)
		return
	}
	fmt.Println("获取返回数据成功")
	mesRes = &message.Message{} 
	err = json.Unmarshal(tf.Buf[:length], &mesRes)
	if err != nil {
		fmt.Println("反序列化返回数据失败")
		return
	}
	return
}