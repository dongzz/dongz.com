package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		fmt.Println("service connect error =", err)
		return
	}
	fmt.Println("server connect success,", conn.RemoteAddr())
	//客户端发送单行数据
	//os.Stdin 代表标准输入[终端]
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取一行用户的输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader err = ", err)
			continue
		}

		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("send line err=", err)
			continue
		}
		fmt.Printf("客户端发送了 %d 字节的数据\n", n)

		//再见line 发送给服务器
		if line == "exit\n" {
			fmt.Println("客户端已退出")
			break
		}
	}
}
