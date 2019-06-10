package main

type StudentMapper interface {
	GetStuBySid(sid string)
}

func main() {
	var studentMapper StudentMapper 
	studentMapper.GetStuBySid("10001")
}