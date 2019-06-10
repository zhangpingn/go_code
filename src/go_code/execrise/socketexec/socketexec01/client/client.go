package main
import (
	"net"
	"fmt"
	"bufio"
	"os"
	"encoding/binary"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.126.1:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端连接服务器异常", err)
		return 
	}

	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取文件异常")
	}
	n, err := conn.Write([]byte(msg))
	var buf []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	fmt.Printf("共写入%d个字节\n", n)
}