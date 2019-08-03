package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(2)

	var num int64 = 0

	var group sync.WaitGroup
	group.Add(2)

	task := make(chan string, 20)

	fmt.Println("start")

	go func() {
		defer group.Done()
		for i := 0; i < 20; i++ {
			//runtime.Gosched()
			//atomic.AddInt64(&num,1)
			task <- fmt.Sprintf("chan num = %v \n", num)
			num++
			fmt.Println("A num =", num)
			msg, ok := <-task
			if ok {
				fmt.Println(msg)
			}
		}
	}()

	go func() {
		defer group.Done()
		for i := 0; i < 20; i++ {
			//runtime.Gosched()
			//atomic.AddInt64(&num,1)
			task <- fmt.Sprintf("chan num = %v \n", num)
			num++
			fmt.Println("B num =", num)
			msg, ok := <-task
			if ok {
				fmt.Println(msg)
			}
		}
	}()

	group.Wait()
	close(task)

	fmt.Println("stop,num=", num)
}
