package message

const (
	LoginMsgType = "LoginMsgType"
	LoginMsgResType = "LoginMsgResType"
)

type Message struct{
	MsgType string `json:"msgType"`
	MsgBody string `json:"msgBody"`
}

type LoginMsg struct{
	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type LoginMsgRes struct{
	ResCode string `json:"resCode"`
	Error error `json:"error"`
}