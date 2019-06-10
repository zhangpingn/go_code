package main
import (
	"fmt"
	"reflect"
)

func test(a, b int) int{
	return a + b
}

func test1(a interface{}, args... interface{}) interface{}{
	val := reflect.ValueOf(a)
	var args1 []reflect.Value = make([]reflect.Value, len(args))
	for k, v := range args {
		args1[k] = reflect.ValueOf(v)
	}
	res := val.Call(args1)
	defer func(){
		res[0] = reflect.ValueOf(50)
		fmt.Println("hello")
	}()
	return getRes(res[0].Interface())
	//return res
}

func getRes(a interface{}) interface{} {
	fmt.Println("hello,world")
	return a
}

func main() {
	var res = test1(test, 10, 20)
	v, ok := res.(int)
	if ok {
		fmt.Println("结果是", v)
	}
}