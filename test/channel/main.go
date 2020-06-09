package main

import (
	"fmt"
	"sync"
)

// channel练习
// 1、启动一个goroutine，生成100个数发送到ch1
// 2、启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3、在main中 从ch2中取值打印出来

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {		// chan<- 这么写表示只能发送（单向通道），确保暴露出去的通道，别人只能做某一项操作，以免形成干扰
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)		// 关闭通道后，可读不可写。对已关闭的通道输出多于写入数量的取值时，不会报错，返回的OK是false，前面那对应的值为0值（int）/false（bool）/''（string）
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	/*for x := range ch1{		// 这样写可可能ch1还没写完就被range了
		ch2 <- x*x
	}*/
	for true {
		x,ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x*x
	}
	once.Do(func() { close(ch2) })	// 确保某个操作只执行一次
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a,b)
	go f2(a,b)
	wg.Wait()
	for ret := range b{
		fmt.Println(ret)
	}

}
