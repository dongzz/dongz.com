package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(2)

	num := 0

	var group sync.WaitGroup
	group.Add(2)

	var mutex sync.Mutex

	fmt.Println("start")

	go func() {
		defer group.Done()
		for i := 0; i < 20; i++ {
			mutex.Lock()
			{
				value := num
				runtime.Gosched()
				value++
				fmt.Println("A")
				num = value
			}
			mutex.Unlock()
		}
	}()

	go func() {
		defer group.Done()
		for i := 0; i < 20; i++ {
			mutex.Lock()
			{
				value := num
				runtime.Gosched()
				value++
				fmt.Println("B")
				num = value
			}
			mutex.Unlock()
		}
	}()

	group.Wait()

	fmt.Println("stop,num=", num)
}
