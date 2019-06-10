package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)
func main(){
	file,err := os.Open("d:/111.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	reader := bufio.NewReader(file)
	var i int = 0
	for{
		str, err := reader.ReadString('\n')
		i++
		if err == io.EOF {
			break
		}
		fmt.Printf("第%v文件内容为%v",i, str)
	}
	fmt.Println("文件读取结束")

	str, err := ioutil.ReadFile("d:/111.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}

	fmt.Println(string(str))
}