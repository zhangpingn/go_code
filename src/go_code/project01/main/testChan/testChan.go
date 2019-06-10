package main
import "fmt"

func divde(a, b int) {
	fmt.Println(a / b)
}
var sum int
func add(b,ch1 chan int, c int) {
	b <- c
	ch1 <- 10
	select {
		case a := <- b:
			fmt.Println("来自通道b的值")
			sum = sum + a
		case a := <- ch1:
			fmt.Println("来自通道ch1的值")
			sum = sum + a
		default :
			fmt.Println("并没有改变值")
	}
	fmt.Println(sum)
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}	
	}()
    divde(10, 1)
    go divde(10, 0)
	var b = make(chan int, 1)
	ch1 := make(chan int, 1)
	//b <- 10
	//a, ok := <- b
	//ch1 <- 10

     add(b, ch1, 10)
	fmt.Println("外部的值",sum)
	 add(b, ch1, 14) 
	fmt.Println("外部的值",sum)
	add(b, ch1, 5)
	fmt.Println("外部的值",sum)
}