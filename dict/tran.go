package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func main() {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var resout bytes.Buffer
	var reserr bytes.Buffer
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

	//[2K[1G
	trans := exec.Command("dict", word)
	trans.Stdout = &resout
	trans.Stderr = &reserr
	trans.Run()

	result := ""
	if reserr.Len() == 0 {
		result = resout.String()
		result = strings.Replace(result, "\x1b[2K", "", -1)
		result = strings.Replace(result, "\x1b[1G", "", -1)
		result = strings.Replace(result, "\n", "", -1)
		result = strings.TrimLeft(result, " ")
	}

	if result == "" {
		return
	}
	command := exec.Command("zenity", "--notification", "--window-icon=\"info\"", "--text="+result)
	command.Run()

	cleanClip := exec.Command("/home/dongzhi/./cleanClip.sh")
	cleanClip.Run()
}
