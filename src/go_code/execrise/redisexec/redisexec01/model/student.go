package model

type Student struct {
	Sid string `json:"sid"`
	Sname string `json:"sname"`
	Age uint `json:"age"`
	Score float64 `json:"score"`
}