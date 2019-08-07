package login

import (
	"dongz.com/ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	fmt.Printf("userId is %v,userPwd is %v\n", userId, userPwd)

	//1,连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client connect server error = ", err)
		return
	}

	defer conn.Close()

	//2. 准备通过conn 发送消息给server
	var msg message.Message
	msg.Type = message.LoginMsgType
	//3. 创建一个LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	//4. 将loginMsg序列化
	marshal, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("loginMsg marshal err = ", err)
		return
	}

	msg.Data = string(marshal)

	//5. 将msg进行序列化
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("msg marshal err = ", err)
		return
	}
	// 到了这里 data 就是我们要发送的数据
	// 6.1 先别data的长度传给server
	// 先获取到data的长度-> 转成[]byte
	pkgLen := uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn Write err = ", err)
		return
	}
	//fmt.Println("客户端发送消息长度成功，n=",n)

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("data Write err = ", err)
		return
	}
	return nil
}
