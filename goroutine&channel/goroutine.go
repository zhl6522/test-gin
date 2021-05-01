package main

import (
	"fmt"
	"strconv"
	"time"
)
//要求主线程和goroutine同时执行

//编写一个函数，每隔一秒输出"Hello World"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() Hello World", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//执行示意图，参考同目录文件：main.go的主线程和协程工作示意图.png
func main() {
	go test()	//开启了一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() Hello GoLand", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
