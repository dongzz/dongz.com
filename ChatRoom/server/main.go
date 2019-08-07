package main

import (
	"dongz.com/ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

func readPkgMsg(conn net.Conn) (msg message.Message, err error) {
	bytes := make([]byte, 1024)

	_, err = conn.Read(bytes[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		err = errors.New("read head err")
		return
	}
	//根据bytes[:4]转成一个uint32
	pgkLen := binary.BigEndian.Uint32(bytes[:4])

	//根据pkgLen读取消息内容
	line, err := conn.Read(bytes[:pgkLen])
	if line != int(pgkLen) || err != nil {
		if err == io.EOF {
			return
		}
		err = errors.New("read message err")
		return
	}

	//反序列化
	//特别注意 &
	err = json.Unmarshal(bytes[:pgkLen], &msg)
	if err != nil {
		err = errors.New("unmarshal err")
		return
	}

	return
}

//处理登录请求
func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
	//1. 先从msg中取出 data，再反序列化成 loginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("unmarshal err")
		return
	}

	//声明一个resultMsg
	var resMsg message.Message
	resMsg.Type = message.LoginResultMsgType

	//声明一个loginResultMsg
	var loginResultMsg message.LoginResultMsg

	//测试 userId = 200 userPwd = 123456
	if loginMsg.UserId == 200 && loginMsg.UserPwd == "123456" {
		fmt.Println("登录成功")
		loginResultMsg.Code = 200
	} else {
		//用户不存在
		loginResultMsg.Code = 500
		loginResultMsg.Error = "id or pwd err"
	}

	//将loginResultMsg 序列化
	data, err := json.Marshal(loginResultMsg)
	if err != nil {
		fmt.Println("loginResultMsg marshal err")
		return
	}

	// 将data赋值给 resMsg
	resMsg.Data = string(data)

	// 对resMsg序列化
	result, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println("loginResultMsg marshal err")
		return
	}

	//发送
	i := len(result)

	return
}

//根据客户端发送来的消息类型，决定调用哪个函数来处理
func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		//处理登录逻辑
		err = serverProcessLogin(conn, msg)
	case message.LoginResultMsgType:

	case message.RegisterMsgType:
		//处理注册
	default:
		fmt.Println("消息类型不存在， 无法处理")
	}
	return
}

//处理和客户端通讯
func process(conn net.Conn) {
	defer conn.Close()

	for {
		msg, err := readPkgMsg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client %v closed \n", conn.RemoteAddr())
				return
			}
			fmt.Println("getMsg err :", err)
			continue
		}

		fmt.Println("msg is ", msg)
	}

}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("server start err = ", err)
		return
	}

	defer listener.Close()

	fmt.Println("server listen 8888 ...")

	//开始监听客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("client connect err = ", err)
			continue
		}
		fmt.Printf("client %v is connect\n", conn.RemoteAddr())

		go process(conn)
	}

}
