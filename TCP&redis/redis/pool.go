package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //数据库的最大连接数，0表示无限制
		IdleTimeout: 100, //最大空闲时间，单位秒
		Dial: func() (redis.Conn, error) { //初始化链接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "benben")
	if err != nil {
		fmt.Printf("conn.Do set=%v\n", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Printf("conn.Do set=%v\n", err)
		return
	}
	fmt.Println("r=", r)
}
