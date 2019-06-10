package message

import (
	"go_code/restudy/netstudy/netstudy03/common/model"
)
const (
	LoginMesType    = "LoginMesType"
	LoginMesResType = "LoginMesResType"
)

type Message struct {
	MsgType string
	Content string
}

type LoginMes struct {
	User model.User
}

type LoginMesRes struct {
	ErrorCode string
	ErrorDesc string
}

