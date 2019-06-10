package main
import "fmt"
func main(){
    defer func(){
        if err := recover() ; err != nil {
            fmt.Println(err)
        }
	}()
    defer func(){
		fmt.Println("three")
        panic("three")
    }()
    defer func(){
		fmt.Println("two")
		panic("two")
    }()
    panic("one")
}