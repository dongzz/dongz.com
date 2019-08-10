package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("xclip", "-out")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	word := ""
	if stderr.Len() == 0 {
		word = stdout.String()
	}

	if word == "" {
		return
	}

	resp, err := http.Get("http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=" + word)

	var tran string
	if err != nil {
		tran = " 查询失败!"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		tran = " 查询失败!"
	}

	result := make(map[string]interface{})

	trim := bytes.Trim(body, " ")

	if err = json.Unmarshal(trim, &result); err != nil {
		tran = " 查询失败!"
	} else {
		if errorCode := result["errorCode"].(float64); errorCode == 0 {
			for _, arr := range result["translateResult"].([]interface{}) {
				for _, value := range arr.([]interface{}) {
					tran += value.(map[string]interface{})["tgt"].(string) + "\n"
				}
			}
		}
	}

	command := exec.Command("zenity", "--notification", "--window-icon=\"info\"", "--text="+tran)
	command.Start()

	cleanClip := exec.Command("/home/dongzhi/./cleanClip.sh")
	cleanClip.Run()
}
