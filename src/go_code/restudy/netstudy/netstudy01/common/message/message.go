package message

const (
	LoginType string = "login"
	LoginResType string = "loginRes"
	OnlineMessageType string = "onlineMessage"
)

type Message struct {
	MesType string
	Data string
}

type LoginMessage struct {
	UserId string
	Pwd string
}

type LoginMessageRes struct {
	Code int
	Desc string
}

type OnlineMessage struct {
	UserId string
}