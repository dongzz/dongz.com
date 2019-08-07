package main

import (
	"dongz.com/ChatRoom/common/login"
	"fmt"
)

var userId int
var userPwd string

func main() {

	//接收用户的选择
	var chechType int
	//判断是否还显示菜单
	flag := true

	for flag {
		fmt.Println("----------  欢迎登录多人聊天系统   --------")
		fmt.Println("----------  1. 登录聊天系统       --------")
		fmt.Println("----------  2. 注册用户          --------")
		fmt.Println("----------  3. 退出系统          --------")
		fmt.Println("----------  请选择（1-3）        --------")
		fmt.Println("----------------------------------------")
		fmt.Scanln(&chechType)

		switch chechType {
		case 1:
			fmt.Println("登录")
			flag = false
		case 2:
			fmt.Println("注册")
			flag = false
		case 3:
			fmt.Println("退出系统")
			flag = false
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}

	//根据用户的选择，显示新的提示信息
	if chechType == 1 {
		fmt.Println("请输入用户的ID")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanln(&userPwd)

		err := login.Login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败 ")
		}
	}
}
