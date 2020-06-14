package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// waitGroup
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1:=rand.Int()			// int64
		r2:=rand.Intn(10)	// 0<=x<10
		fmt.Println(r1,r2,0-r1,0-r2)
	}
}
func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

func a() {
	defer wg.Done()
	for i :=0 ; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i :=0 ; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
var wg sync.WaitGroup	// 多个goroutine等待组
// 程序启动之后会创建一个主goroutine去执行
func main() {

	/*for i := 0; i < 100; i++ {
		//go hello(i)		// 开启一个单独的goroutine去执行hello函数（任务）
		go func(i int) {
			fmt.Println(i)	// 用的是函数参数的那个i，不是外面的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second*3)*/
	// main函数结束了 由main函数启动的goroutine也都结束了

	//f()
	/*for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// ?如何知道这10个goroutine都结束了
	wg.Wait()*/	// 等待wg的计算器减为0

	wg.Add(2)
	runtime.GOMAXPROCS(4)		// 默认CPU的逻辑核心数（go1.5之后），默认跑满整个CPU。在mac上可以看到 a/b的数据交叉执行
	//runtime.GOMAXPROCS(1)		// 只占用一个核，比如日志收集。
	go a()
	go b()
	wg.Wait()
	fmt.Println(runtime.NumCPU())
	// goroutine调度模型(goroutine的本质)
	/*GMP
	G 就是个goroutine
	M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系
	P管理着一组goroutine队列(管理者)*/

	/*Go语言中的操作系统线程和goroutine的关系：
		一个操作系统线程对应用户态多个goroutine。
		go程序可以同时使用多个操作系统线程。
		goroutine和OS线程是多对多的关系，即m:n。*/
	// M:N 把m个goroutine分配给n个操作系统线程去执行。
	// goroutine初始栈的大小是2k。

	// channel
	/*
	为什么需要channel？
		通过channel实现多个goroutine之间的通信。
		csp：通过通信来共享内存。
	channel是一种类型，一种引用类型。make函数初始化之后才能使用。（slice、map、channel）
		channel声明 var ch chan 元素类型
		channel初始化 ch=make(chan 元素类型，[缓冲区大小])
		channel的操作：
			发送：ch <- 100
			接收：x := <-ch
			关闭：close(ch)	非必须
	带缓冲区的通道与不带缓冲区的通道
		快递员送快递的示例，有缓冲区就是有快递柜。
	从通道中取值：
		for true {
			x,ok := <-ch1
			if !ok {				// 什么时候ok=false？	ch1通道被关闭的时候
				break
			}
			ch2 <- x*x
		}
		for ret := range b{		// 什么时候for range会退出？	b通道被关闭的时候
			fmt.Println(ret)
		}
	单向通道：
		通常是用作函数的参数，只读通道 <-chan int 和只写通道 chan<- int
	通道的各种考虑情况：https://www.liwenzhou.com/posts/Go/14_concurrence/里的通道总结
	select多路复用
		同一时刻有多个通道要操作的场景下，使用select。
		func main() {
			ch := make(chan int, 1)
			for i := 0; i < 10; i++ {
				select {
				case x := <-ch:
					fmt.Println(x)
				case ch <- i:
				}
			}
		}
		使用select语句能提高代码的可读性。
			可处理一个或多个channel的发送/接收操作。
			如果多个case同时满足，select会随机选择一个。
			对于没有case的select{}会一直等待，可用于阻塞main函数。

	*/
}
