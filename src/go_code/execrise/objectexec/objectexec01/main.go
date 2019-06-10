package main
import (
	"fmt"
)

type Person struct{
	Name string
	Age uint
}

func (p Person) test(){
	fmt.Println("hello,world")
}
func (p Person) String() string{
	str := fmt.Sprintf("person的name是%v, 年龄是%v", p.Name, p.Age)
	return str
} 
func main() {
	var p *Person = new(Person)
	(*p).test()
	p.Name = "zhangsan"
	p.Age = 10
	fmt.Println(p)
	var p1 = &Person{"lisi", 20}
	p1 = &Person{
		Name : "wangwu",
		Age : 30,
	}
	fmt.Println(p1)
}