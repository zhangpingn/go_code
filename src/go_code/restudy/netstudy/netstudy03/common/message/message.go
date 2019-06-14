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
	Blood uint
	HeroName string
	Power uint
	ErrorCode uint
	ErrorDesc string
}

type AttackMes struct {
	FromHeroName string
	ToHeroName string
	Power uint
}

