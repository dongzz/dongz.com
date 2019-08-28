package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

type Langdetect struct {
	ErrorNum int    `json:"error"`
	Msg      string `json:"msg"`
	Lan      string `json:"lan"`
}

var appid = "20190818000327515"
var appSecure = ""

func getWord() (word string, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("xclip", "-out")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	if stderr.Len() == 0 {
		word = stdout.String()
	}

	if word == "" {
		err = errors.New("无选中文本")
	}
	return
}

func checkType(word string) (wordType, toType string, err error) {
	var body []byte
	resp, err := http.PostForm("https://fanyi.baidu.com/langdetect",
		url.Values{"query": {word}})
	if err != nil {
		err = errors.New("文本语种判断失败")
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("文本语种获取失败")
		return
	}

	var lan Langdetect
	toType = "zh"
	//提取参数
	err = json.Unmarshal(body, &lan)

	if err != nil {
		return
	}

	if !(lan.ErrorNum == 0 && lan.Msg == "success") {
		err = errors.New("无文本语种")
		return
	}

	wordType = lan.Lan

	if wordType == "zh" {
		toType = "en"
	}

	return
}

func createSalt() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

func getSign(appid, q, salt, secure string) string {
	data := []byte(appid + q + salt + secure)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func main() {
	//获取选中文字
	word, err := getWord()
	if err != nil {
		return
	}

	//判断文字类型
	wordType, toType, err := checkType(word)
	if err != nil {
		return
	}

	//调用百度API接口
	salt := createSalt()
	sign := getSign(appid, word, salt, appSecure)

	var body []byte

	resp, err := http.Get("http://api.fanyi.baidu.com/api/trans/vip/translate?q=" + word + "&from=" + wordType + "&to=" + toType + "&appid=" + appid + "&salt=" + salt + "&sign=" + sign)

	if err != nil {
		word = "查询失败"
		goto zeni
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

zeni:
	//提示
	command := exec.Command("zenity", "--notification", "--window-icon=\"info\"", "--text="+word)
	command.Start()
}
