package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("服务器开始监听！")

	//tcp 表示使用的网络协议
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("server error=", err)
		return
	}
	//延时关闭
	defer listener.Close()

	fmt.Println("等待客户端连接！")
	//循环，待客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			continue
		}

		//启动协程为客户端服务
		fmt.Println("client connect success,", conn.RemoteAddr())

		go func(conn net.Conn) {
			defer conn.Close()
			//循环接受客户端信息
			for {
				bytes := make([]byte, 1024)
				n, err := conn.Read(bytes)
				if err != nil {
					if err == io.EOF {
						continue
					}
					fmt.Println("message read err", err)
					break
				}

				line := string(bytes[:n])
				if line == "exit\n" {
					fmt.Printf("客户端 %v 已退出\n", conn.RemoteAddr())
					break
				}

				fmt.Print("message is: ", line)
			}

		}(conn)
	}
}
