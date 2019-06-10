package main

import (
	"html"
	"fmt"
)
func test(a, b int) (sum int) {
	sum = a + b
	return
}
func main() {
	s := html.UnescapeString("&amp;")
	s = html.EscapeString(s)
	fmt.Println(s)
}