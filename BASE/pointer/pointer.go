package main
import (
	"fmt"
	"flag"
)
var mode = flag.String("mode","","process mode")
func main(){
	//解析命令行参数
	flag.Parse()

	//输出命令行参数
	fmt.Println(*mode)

	x, y := 1, 2
	fmt.Println(x,y)
	x,y = swap1(x,y)
	fmt.Println(x,y)
	swap2(&x,&y)
	fmt.Println(x,y)
	swap3(&x,&y)
	fmt.Println(x,y)
}

func swap1(x,y int) (int,int) {
	x,y = y,x
	return x,y
}

func swap2(x,y *int){
	t := *x
	*x = *y
	*y = t
}
func swap3(x,y *int){
	x,y = y,x
}
