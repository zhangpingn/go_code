package main
import (
	"fmt"
	"runtime"
	_"time"
)

func test(i int){
	fmt.Println("hello, world", i)
}

func main() {
	var cpuNum = runtime.NumCPU()
	fmt.Printf("cpu的个数是%v\n", cpuNum)
	runtime.GOMAXPROCS(5)
	for i:=0; i<10; i++ {
		go test(i)
	}
	for i:=0; i<10; i++ {
		fmt.Println("hello, golang", i)
	}
	//time.Sleep(time.Second)
}