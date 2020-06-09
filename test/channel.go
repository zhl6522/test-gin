package main

import (
	"fmt"
	"sync"
)

var a []int		// slice
var b chan int	// 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)		// nil
	b = make(chan int)	// 不带缓存区通道的初始化，通道必须使用make函数初始化才能使用！
	//b <- 10				// 死锁:fatal error: all goroutines are asleep - deadlock!
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b了...")
	wg.Wait()
}
func bufChannel() {
	fmt.Println(b)		// nil
	b = make(chan int, 10)	// 如果值比较大，可以使用指针
	b <- 10
	fmt.Println("10发送到通道b了...")
	b <- 20
	fmt.Println("20发送到通道b了...")
	x := <- b
	fmt.Println("从通道b中取到了", x)
	close(b)
}
func main() {

	bufChannel()
	// 通道的操作
	//1、发送：ch1 <- 1
	//2、接收：<- ch1
	//3、关闭：close()
}
