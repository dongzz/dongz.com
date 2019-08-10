package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println(user.Current())
	//mkdir := os.remove("/home/dongzhi/Desktop/123", os.ModeDir)
	//fmt.Println(mkdir)

	file, _ := os.Open("/home/dongzhi/Desktop/img/打雪仗.png")
	info, _ := file.Stat()
	fmt.Println(info)
}
