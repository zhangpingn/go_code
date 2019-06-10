package main
import (
	"fmt"	
	"io/ioutil"
)

func main() {
	filePath := "d:/111.txt"
	comment, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}
	fmt.Println(string(comment))
}