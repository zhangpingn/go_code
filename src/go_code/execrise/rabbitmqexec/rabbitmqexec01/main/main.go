package main

import (
	_ "github.com/streadway/amqp"
	"go_code/execrise/rabbitmqexec/rabbitmqexec01/utils"
)

func main(){
	conn := utils.Conn
	channel, _ := conn.Channel()
	utils.PublishMsg(channel)
}