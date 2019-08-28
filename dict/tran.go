package main

import (
	"bytes"
	"os/exec"
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

	trans := exec.Command("fy", word)
	trans.Stdout = &resout
	trans.Stderr = &reserr
	trans.Run()

	result := ""
	if reserr.Len() == 3 && reserr.String() == "- \n" {
		result = resout.String()
	}

	if result == "" {
		return
	}
	command := exec.Command("zenity", "--notification", "--window-icon=\"info\"", "--text="+result)
	command.Run()

	cleanClip := exec.Command("~/./cleanClip.sh")
	cleanClip.Run()
}
