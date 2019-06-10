package main

import (
	"net"
	"fmt"
	"go_code/restudy/netstudy/netstudy01/server/utils"
)

func main() {
	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println("创建监听失败", err)
		return 
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("获取连接失败", err)
			continue
		}
		go utils.Process(conn)
	}
}