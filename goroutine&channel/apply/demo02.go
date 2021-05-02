package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world")
		time.Sleep(time.Second)
	}
}
func test() {
	//goroutine中使用recover，解决协程中出现panic导致程序崩溃的问题，不影响主线程和其他协程
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err :=recover(); err != nil {
			fmt.Println("test()协程发生错误", err)
		}
	}()
	//定义一个map
	var myMap map[int]string
	myMap[0] = "Goland"	//panic: assignment to entry in nil map
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)
	}
}
