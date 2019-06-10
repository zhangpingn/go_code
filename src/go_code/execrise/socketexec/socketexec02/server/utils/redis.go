package utils

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"fmt"
	"go_code/execrise/socketexec/socketexec02/common/entity"
	"encoding/json"
)

var Conn redis.Conn

func init(){
	conn, err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		os.Exit(1)
	}
	Conn = conn
}

func GetUserById(id string)(user entity.User) {
	userMsg, err := redis.String(Conn.Do("HGet", "users", id))
	if err != nil {
		fmt.Printf("获取%s用户数据失败\n", id)
		return
	}
	//fmt.Printf("userMsg的值是%v,类型是%T\n", userMsg,userMsg)
	//user = entity.User{}
	//var stuMsg = "{\"sid\":\"10001\",\"sname\":\"zhangsan\"}"
	//stu = &Student{}
	err = json.Unmarshal([]byte(userMsg), user)
	//err = json.Unmarshal(userMsg, user)
	return
}