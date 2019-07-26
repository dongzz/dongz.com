package main
import "fmt"

func main(){
	var map1 map[string]int
	fmt.Printf("map1's type is %T,value is %v\n",map1,map1)

	map1 = map[string]int{"one":1,"two":2}
	map2 := map1
	map2["two"] = 3

	fmt.Printf("map2['two'] is %v,map['two'] is %v\n",map2["two"],map1["two"])
}
