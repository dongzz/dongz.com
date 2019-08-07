package message

const (
	LoginMsgType       = "LoginMsg"
	LoginResultMsgType = "LoginResultMsg"
	RegisterMsgType    = "RegisterMsg"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResultMsg struct {
	Code  int    `json:"code"`  //返回状态码 500 表示该用户未注册  200 表示登录成功
	Error string `json:"error"` //返回错误信息
}

type registerMsg struct {
}
