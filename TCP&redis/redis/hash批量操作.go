package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1、链接到redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("redis.Dial err=%v\n", err)
		return
	}

	defer c.Close()
	//2、通过go向redis写入数据 string [key-val]
	_, err = c.Do("HMset", "user02", "name", "mumu", "age", 17)
	if err != nil {
		fmt.Printf("redis.Do HMset err=%v\n", err)
		return
	}
	//3、通过go向redis读取数据 string [key-val]
	//因为返回r是interface{}
	//因为name对应的值是string，因此我们需要转换
	//nameString := r.(string)	//不可以断言：panic: interface conversion: interface {} is []uint8, not string
	//使用redis自带的方法 redis.String()处理
	r, err := redis.Strings(c.Do("HMget", "user02", "name", "age"))
	if err != nil {
		fmt.Printf("redis.Do HMget err=%v\n", err)
		return
	}
	for k, i := range r {
		fmt.Printf("r[%d]=%v\n", k, i)
	}

}
