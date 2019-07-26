package main

import (
	"fmt"
	"strings"
	"errors"
)

type point struct {
	x,y int
}

func main(){
	str := "cute girl"
	flo := 12.658
	fmt.Printf("%v \n",str)
	fmt.Printf("%q \n",str)
	fmt.Printf("%f \n",flo)
	fmt.Printf("%e \n",flo)
	
	p := point {2,4}
	fmt.Printf("%T , %v \n",p,p)
	
	str2 := fmt.Sprintf("%b",123)
	fmt.Printf("%T , %v \n",str2,str2)

	fmt.Printf("%v \n",30*60*60)


	num := 6
	j:= num
	for i:=0;i<=num;i++ {
		fmt.Printf("%v + %v = %v \n",i,j,i+j)
		j--
	}

	count,sum:= 0,0
	for i:=0;i<100;i++{
		if i%9==0 {
			count++
			sum+=i
		}
	}
	fmt.Printf("count = %v,sum = %v \n",count,sum)

	for i:=0;i<6;i++ {
		for j:=0;j<6-i;j++{
			fmt.Print(" ")
		}
		for k:=0;k<2*i+1;k++{
			fmt.Print("*")
		}
		fmt.Println("")
	}

	night()
	//斐波那契数列
	fmt.Println("请输入要计算的斐波那契数:")
	fib := 10
	fmt.Scanln(&fib)
	fmt.Printf("fibonacci: f(%v)=%v \n",fib,fibonacci(fib))

	//f(1)=3;f(n)=2*f(n-1)+1;
	fmt.Printf("fn: f(%v)=%v \n",fib,fn(fib))

	//可变参数
	fmt.Printf("sum():%v\n",Sum(1,2,5,6,1,2,11))
	//闭包
	f1 := addUpper()
	fmt.Println("f1 1:",f1(1))
	fmt.Println("f1 2:",f1(2))
	f2 := addUpper()
	fmt.Println("f2 1:",f2(2))

	//闭包
	f3:=makeSuffix(".jpg")
	fmt.Println("new name of 123",f3("123"))

	//错误处理
	errorTest()
	fmt.Println("错误处理以后")
}

func errorTest(){
	defer func (){
		err := recover()
		if err!=nil {
			fmt.Println("err=",err)
		}
	}()
	panic(errors.New("测试panic"))
	a:=10
	b:=0
	result:=a/b
	fmt.Println("a/b=",result)
}

func makeSuffix(sub string) func(string) string{
	return func(name string) string{
		if strings.HasSuffix(name,sub) {
			return name
		} else {
			return name+sub
		}
	}
}

func addUpper() func(int) int{
	n:= 10
	return func (x int) int {
		n = n+x
		return n
	}
}
func Sum(n int,args... int) (sum int) {
	sum = n
	for i:=0;i<len(args);i++{
		sum += args[i]
	}
	return
}

func fn(n int) int{
	if n <=1 {
		return 3
	}
	return 2*fn(n-1)+1
}

func fibonacci(n int) int{
	if n <=2  {
		return 1 
	}
	return fibonacci(n-1)+fibonacci(n-2)
}

func night() {
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v * %v = %v   ",j,i,j*i)
		}
		fmt.Println("")
	}
}
