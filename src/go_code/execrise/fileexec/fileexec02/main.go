package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	filePath := "d:/111.txt"
	file, err := os.OpenFile(filePath, os.O_RDONLY, 7777)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return 
	}
	reader := bufio.NewReader(file)
	for {
		comment, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件已被读完")
			break
		}
		fmt.Println(comment)
	}
}