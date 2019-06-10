package utils

import (
	"github.com/streadway/amqp"
	"fmt"
)

var (
	Conn *amqp.Connection
)

func init(){
	Conn, _ = amqp.Dial("amqp://guest:guest@10.7.25.37:5672/")
}
func PublishMsg(channel *amqp.Channel){
	// var msg string = "hello, world";
	// channel.Publish("exchange.direct", "h", false, false, amqp.Publishing{
	// 	ContentType: "text/plain",
	// 	Body: []byte(msg),
	// })

	resMsg, _, _ := channel.Get("hello", true)
	fmt.Println(string(resMsg.Body))
}