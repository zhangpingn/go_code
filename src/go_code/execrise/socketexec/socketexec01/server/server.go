package main

import (
	"net"
	"fmt"
)

func process(conn net.Conn){
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("服务器端获取客户端数据异常", err)
		return
	}
	fmt.Print(string(buf[:n]))
	//fmt.Print(buf)
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器监听服务异常", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器获取连接异常", err)
			continue
		}
		//fmt.Printf("服务器获取连接为%v\n", conn)
		go process(conn)
	}
}