package main

import (
	"fmt"
	"flag"
)

func main() {
	var username string
	var pwd string
	var port uint
	var ip string

	flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.UintVar(&port, "port", 3306, "端口，默认是3306")
	flag.StringVar(&ip, "ip", "localhost", "ip,默认是localhost")
	flag.Parse()
	fmt.Printf("username is %v, password is %v, port is %v, ip is %v", username, pwd, port, ip)
	
}