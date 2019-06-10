package message
import (
	"go_code/restudy/netstudy/netstudy02/common/model"
)

const (
	LoginMesType    	 = "LoginMesType"
	LoginMesResType		 = "LoginMesResType"
	RegisterMesType 	 = "RegisterMesType"
	RegisterMesResType 	 = "RegisterMesResType"
	OnlineMesType        = "OnlineMesType"
	InfoMesType          = "InfoMesType"
	InfoMesResType       = "InfoMesResType"
)

type Message struct {
	MsgType string
	Data string
}

type LoginMes struct {
	User model.User
}

type LoginMesRes struct {
	ErrorCode int
	ErrorDesc string
	Users []model.User
}

type RegisterMes struct {
	User model.User
}

type RegisterMesRes struct {
	ErrorCode int
	ErrorDesc string
}

type OnlineMes struct {
	User model.User
}

type InfoMes struct {
	FromUserId string
	ToUserId string
	Info string
}

type InfoMesRes struct {
	ErrorCode int
	ErrorDesc string
}