package main

import "fmt"

/**
稀疏数组
目的，节约多维数组的保存空间，节约内存

需求，将一个五子棋盘数据存盘，同时支持复盘
 0  0  0  0  0  0  0  0  0  0  0
 0  0  0  0  0  0  0  0  0  0  0
 0  0  0  1  0  0  0  0  0  0  0
 0  0  1  2  2  0  0  0  0  0  0
 0  0  0  1  2  0  0  0  0  0  0
 0  0  0  1  2  0  0  0  0  0  0
 0  0  0  0  1  0  0  0  0  0  0
 0  0  0  0  0  0  0  0  0  0  0
 0  0  0  0  0  0  0  0  0  0  0
 0  0  0  0  0  0  0  0  0  0  0
 0  0  0  0  0  0  0  0  0  0  0
*/

type sparse struct {
	row   int
	col   int
	value int
}

//定义如上五子棋盘
//1为白子，2为黑子
func main() {
	gomoku := [11][11]int{
		2: {3: 1},
		3: {2: 1, 3: 2, 4: 2},
		4: {3: 1, 4: 2},
		5: {3: 1, 4: 2},
		6: {4: 1},
	}

	//定义一个切片，记录值
	var sparseArray []sparse

	//标准稀疏数组需要  记录原有二维数组规模
	//sparseArray = append(sparseArray,sparse{len(gomoku), len(gomoku[0]), 0})

	//打印
	//同时将值存入切片(存盘）
	for i, arr := range gomoku {
		for j, value := range arr {
			fmt.Printf("%d\t", value)
			if value > 0 {
				sparseArray = append(sparseArray, sparse{i, j, value})
			}
		}
		fmt.Println()
	}

	for i, arr := range sparseArray {
		fmt.Printf("index: %d,row: %d,col: %d,value: %d \n", i, arr.row, arr.col, arr.value)
	}

	//复盘
	var newGomoku [11][11]int

	for _, arr := range sparseArray {
		newGomoku[arr.row][arr.col] = arr.value
	}

	//打印
	for _, arr := range newGomoku {
		for _, value := range arr {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
	fmt.Println(gomoku == newGomoku)
}
