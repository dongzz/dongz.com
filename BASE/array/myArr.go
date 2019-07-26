package main
import (
	"fmt"
)

func main(){
	var arr [5]int=[5]int{1,2,3,4,5}

	fmt.Println(arr)


	var sli []int = arr[3:4]
	fmt.Printf("sli value is :%v\n sli lenth is :%v\n sli cap is : %v \n",sli,len(sli),cap(sli))

	sli[0]=8
	sli = sli[0:len(sli)+1]
	fmt.Println(sli)

	fmt.Println(arr)

	var sli2 = new([]int)
	fmt.Printf("sli2's type is %T,sli2's value is %v\n",sli2,sli2)
	var sli3 = make([]int,5)
	fmt.Printf("sli3's type is %T,sli3's value is %v\n",sli3,sli3)

	copy(sli3,sli)
	fmt.Printf("sli3's value is %v\n",sli3)
	sli3 = append(sli3,4,5,6)
	fmt.Printf("sli3's value is %v\n",sli3)
}
