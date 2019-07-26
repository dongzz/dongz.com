package main
import (
	"fmt"
)

func main(){
	fmt.Println("请输入一个整数：")
	var num int
	fmt.Scanln(&num)
	result :=fibonacci()
	for i:=1;i<num;i++ {
		result(i);
	}
	fmt.Printf("fibonacci(%v)=%v\n",num,result(num))
}

func fibonacci() func(int) int{
        sum:=1
	sum2:=1
	return func(num int) int{
		if(num>2){
			temp:=sum
			sum=sum2
			sum2=temp+sum2
		}
		return sum2;
	}
}
