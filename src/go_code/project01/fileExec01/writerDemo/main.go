package main

import (
	"fmt"
	_ "bufio"
	_ "io/ioutil"
	_ "os"
	"encoding/json"
)

func main() {
	// file, err := os.Open("d:/111.txt")
	// if err != nil {
	// 	fmt.Println("打开文件失败,", err)
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(file)
	// writer.WriteString("hello")
	str := "{\"address\":\"少林寺\", \"age\":30}"
	var a map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(str), &a)
	fmt.Println(a)
}