package main
import "fmt"

func main() {
	var a, b int = 100, 200
	fmt.Println(&a, &b)
	change(&a, &b)
	fmt.Println(&a, &b)
}

func change(a * int, b * int){
	*a, *b = *b, *a
}