package main
import "fmt"

type Person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int//指针
	slice []int//切片
	map1 map[string]string //map
}

func main(){

	//定义结构体变量
	var me Person
	me.Name = "dongzhi"
	me.Age = 27
	me.Scores[1] = 10

	me.ptr = &me.Age 
	me.slice = make([]int,10)
	me.slice[4] = 200

	me.map1 = make(map[string]string)
	me.map1["key1"] = "232"

	fmt.Printf("me is %v\n",me)
	fmt.Printf("me type is %T\n",me)

	//new
	you := new(Person)
	fmt.Printf("you type is %T \n",you)
	you.Name = "李四"

	it := you
	it.Name = "王五"
	fmt.Printf("you value is %v\n,it value is %v\n",you,it)

	there := me
	there.Name = "张三"
	fmt.Printf("me value is %v\n,there value is %v\n",me,there)
}
