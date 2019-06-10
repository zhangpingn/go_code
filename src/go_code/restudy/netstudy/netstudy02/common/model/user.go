package model

type User struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	UserStatus string `json:"userStatus"`
	Password string `json:"password"`
}