package main

import (
	bytes2 "bytes"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

type My struct {
	*Person
	position string
}

type man interface {
	name()
}

func (my My) name() {
	my.getName()
}

func (p *Person) getName() {
	fmt.Println(p.Name)
}

type myInt int

func change(a, b *myInt) {
	*a, *b = *b, *a
}

func main() {

	//定义结构体变量
	var me Person
	me.Name = "dongzhi"
	me.Age = 27
	me.Scores[1] = 10

	me.ptr = &me.Age
	me.slice = make([]int, 10)
	me.slice[4] = 200

	me.map1 = make(map[string]string)
	me.map1["key1"] = "232"

	fmt.Printf("me is %v\n", me)
	fmt.Printf("me type is %T\n", me)

	//new
	you := new(Person)
	fmt.Printf("you type is %T \n", you)
	you.Name = "李四"

	it := you
	it.Name = "王五"
	fmt.Printf("you value is %v\n,it value is %v\n", you, it)

	there := me
	there.Name = "张三"
	fmt.Printf("me value is %v\n,there value is %v\n", me, there)

	m := myInt(1)
	n := myInt(3)
	fmt.Printf("m: %v,n: %v \n", m, n)

	change(&m, &n)
	fmt.Printf("m: %v,n: %v \n", m, n)

	fmt.Println("curl")
	var bytes bytes2.Buffer
	bytes.Write([]byte("hello\n"))

	io.Copy(os.Stdout, &bytes)

	my := My{
		Person: &Person{
			Name:   "dongzhi",
			Age:    27,
			Scores: [5]float64{3: 0.2},
			ptr:    nil,
		},
		position: "",
	}

	var dong man
	dong = my
	dong.name()

	my.getName()
}
