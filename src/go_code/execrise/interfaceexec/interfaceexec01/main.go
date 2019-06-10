package main
import (
	"fmt"
)
type Usb interface{
	start()
	end()
}

type Computer struct{

}

func (c Computer) Working(usb Usb){
	usb.start()
	usb.end()
}

type Phone struct{

}

func (phone Phone) start(){
	fmt.Println("phone is starting...")
} 
func (phone Phone) end(){
	fmt.Println("phone is ending...")
}
func main() {
	var c Computer
	c.Working(Phone{})

	// var b interface{} = c
	// var c1 Computer 
	// c1 = Computer(b)
	var i1 int = 10
	var i2 interface{} 
	i2 = i1
	var i3 int = i2.(int)
	fmt.Println(i3)
}