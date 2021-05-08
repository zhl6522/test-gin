package main

import (
	"fmt"
	"time"
)
//求80000以内的素数
func writeData(putChan chan int) {
	for i := 1; i < 80000; i++ {
		putChan <- i
	}
	close(putChan)
}
func readData(putChan chan int, primeChan chan int, exitChan chan bool) {
	for true {
		v, ok := <-putChan
		if !ok {
			break
		}
		prime := 1
		if v == 1 {
			prime = 0
			break
		}
		put := v / 2
		for i := 2; i <= put; i++ {
			if v%i == 0 {
				prime = 0
				break
			}
		}
		if prime == 1 {
			primeChan <- v
		}
	}
	fmt.Println("有一个readData协程因为取不到数据，退出")
	exitChan <- true
}

func send(chan1 chan<- int) {}
func recv(chan1 <-chan int) {}

func main() {
	/*
	var chan1 chan int	//默认情况下，管道是双向的 可读可写
	go send(chan1)	//设置发送只写
	go recv(chan1)	//设置接收只读
	*/
	/*//声明为只写，读取会报错
	var chan2 chan<- int
	chan2 = make(chan int, 3)*/
	//var chan3 <-chan int	//声明为只读

	exitChan := make(chan bool, 4)
	putChan := make(chan int, 20000)
	primeChan := make(chan int, 20000)
	start := time.Now().UnixNano()
	go writeData(putChan)
	for i := 0; i < 4; i++ {
		go readData(putChan, primeChan, exitChan)
	}
	//这里我们主线程，进行处理
	go func() { //起一个协程去处理（匿名函数）避免阻塞在那里
		for i := 0; i < 4; i++ {	//由于当前电脑4核开四个协程基本就实现了最大化，开多了也没多大变化
			<-exitChan //这里取不到就会一直等待
		}
		end := time.Now().UnixNano()
		fmt.Println("使用协程耗时=", end-start)	//协程是普通方法速度的两倍（协程数） 参考文件：apply2.test.go
		//当我们从exitChan取出4个结果，就可以放心关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan，把结果取出
	for true {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数值=%d\n", res)
	}
	fmt.Println("main() 线程结束")
	// 结论：
}
