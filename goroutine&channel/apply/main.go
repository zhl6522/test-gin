package main

import (
	"fmt"
	_ "time"
)

func writeData(intChan chan int) {
	for i := 0; i < 60; i++ {
		intChan<-i
		fmt.Printf("writeData：%v\n", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for true {
		//time.Sleep(time.Second)	//只要有读取操作 就不会发生deadlock 会阻塞等待
		v,ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData取出的内容：%v\n", v)
	}
	//readData读取完数据后，即任务完成
	exitChan<-true
	close(exitChan)
}
/*
应用实例
完成goroutine和channel协同工作的案例，要求如下：
1、开启一个writeData协程，向管道intChan中写入50个整数。
2、开启一个readData协程，从管道intChan中读取writeData写入的数据。
3、注意：writeData和readData操作的是同一个通道。
4、主线程需要等待writeData和readData协程都完成工作才能退出。【管道】
*/
func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)	//如果本行注释（管道只有写没有读） 47行会发生管道未关闭（chan receive），10行也会死锁
	for true {
		_,ok := <-exitChan
		if !ok {
			break
		}
	}
}
