package main

import (
   fmt1	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt1.Println(r.URL.Path)
	fmt1.Fprintf(w, "hello,world")
}
func init() {
	fmt1.Println("你好啊")
}
func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":18892", nil)
	if err != nil {
		fmt1.Println(err)
	}

}