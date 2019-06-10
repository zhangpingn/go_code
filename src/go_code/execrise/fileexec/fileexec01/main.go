package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	filePath := "D:/111.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 7777)
	if err != nil {
		fmt.Println("文件打开失败！", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, World\r\n")
	writer.Flush()
}