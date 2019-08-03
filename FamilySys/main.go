package main

import "fmt"

func main() {

	var flag = true
	var sli []int

	for flag {
		fmt.Println("请输入一个数:")

		var index int
		fmt.Scanln(&index)

		if index == -1 {
			flag = false
		}
		sli = append(sli, index)
	}

	fmt.Println(sli)
}
