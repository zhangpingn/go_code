package utils

import (
	"go_code/execrise/redisexec/redisexec01/model"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
)

func SetStudent(sid string,ptr *string) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("redis连接失败")
		return err
	}
	_, err = conn.Do("set", sid, *ptr)
	if err != nil {
		fmt.Println("set数据失败")
	}
	return err
}

func GetStudent(sid string) model.Student {
	var stu model.Student
	conn, err := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("redis连接失败")
		return stu
	}
	stuStr, err := redis.String(conn.Do("get", sid))
	if err != nil {
		fmt.Println("get数据失败")
	}
	json.Unmarshal([]byte(stuStr), &stu)
	return stu
}

func HSetStudent(sid string,stu *model.Student) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("redis连接失败")
		return err
	}
	stuStr, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json序列化stu失败", err)
		return err
	}
	fmt.Printf("stuStr的类型是%T", string(stuStr))
	_, err = conn.Do("HSet", "stu", sid, string(stuStr))
	if err != nil {
		fmt.Println("HSet失败")
		return err
	}
	return nil
}
func HGetStudent(sid string) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("redis连接失败")
		return err
	}
	res, err := conn.Do("HGetall", sid)
	if err != nil {
		fmt.Println("HGet失败")
		return err
	}
	fmt.Println(res)
	return nil
}