package main
import "fmt"

var a, b, c int
const d int = 5
func hello() {
	a = 1
	b = 2
	c = 3
	fmt.Printf("a=%v b=%v c=%v",a, b, c)
}

func main() {
	hello()
}