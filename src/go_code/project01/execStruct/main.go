package main
import (
	"fmt"
)
type fruit struct {
	name string
	price float32
	left, right *fruit
}

func (node *fruit) get() string {
	fmt.Printf("%p\n", node)
	return node.name
}

func (node *fruit) set(name string) {    
	fmt.Printf("hello%p\n", node)
	node.name = name
}
func main() {
	var leftFruit fruit
	leftFruit.name = "香蕉"
	leftFruit.price = 13.5
	apple := fruit{name : "苹果", price : 12}
	apple.price = 13.4
	apple.left = &leftFruit
	fmt.Println(apple.name)
	fmt.Println(apple.price)
	fmt.Println(apple.left.name)
	fmt.Printf("%p\n", &apple)
	apple.set("张三")
	fmt.Println(apple.get())
	fmt.Printf("%p", &apple)
}