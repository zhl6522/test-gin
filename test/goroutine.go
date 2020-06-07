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
var wg sync.WaitGroup	// 多个goroutine同步
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
	runtime.GOMAXPROCS(4)		// 默认CPU的逻辑核心数，默认跑满整个CPU。在mac上可以看到 a/b的数据交叉执行
	go a()
	go b()
	wg.Wait()
	fmt.Println(runtime.NumCPU())
	// goroutine调度模型
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

}
