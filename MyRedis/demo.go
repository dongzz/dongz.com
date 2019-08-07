package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1,连接redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	defer conn.Close()

	//2，通过go向redis中写入key-value
	_, err = conn.Do("Set", "name", "dongz")
	if err != nil {
		fmt.Println("set err=", err)
	}
	reply, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	fmt.Printf("name is %v\n", reply)
}
