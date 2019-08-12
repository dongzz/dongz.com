package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
	Age  int
}

func main() {
	my := People{"董志", 27}

	of := reflect.TypeOf(my)
	fmt.Println(of.String())
	fmt.Println(of.FieldByIndex([]int{1}))
	fmt.Println(of.FieldByName("Age"))

	valueOf := reflect.ValueOf(my)
	fmt.Println(valueOf.FieldByName("Age"))
	fmt.Println(valueOf.String())

	fmt.Println(valueOf.CanSet())

	elem := reflect.ValueOf(&my).Elem()
	fmt.Println(elem.CanSet())

	elem.FieldByName("Age").SetInt(28)
	fmt.Println(my)
}
