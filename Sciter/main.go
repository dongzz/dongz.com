package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	//加载文件
	w.LoadFile("demo1.html")
	//设置标题
	w.SetTitle("你好，世界")
	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}
