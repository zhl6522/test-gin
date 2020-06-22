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
var notifyCh = make(chan struct{}, 5)		// struct{}不占用空间，int占用八个字节，所以更节省空间。通常用来做通知用。
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	/*for x := range ch1{		// 这样写可可能ch1还没写完就被range了
		ch2 <- x*x
	}

	notifyCh <- struct{}{}
	// type cat struct {}	// 声明类型
	// var c1 = cat{}		// 实例化
	*/
	for true {
		x,ok := <-ch1
		if !ok {				// 什么时候ok=false？	ch1通道被关闭的时候
			break
		}
		ch2 <- x*x
	}
	f := func() {
		close(ch2)
	}
	once.Do(f)	// 确保某个操作只执行一次	Do只能接收一个没有参数没有返回值的函数（匿名/闭包）
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a,b)
	go f2(a,b)
	wg.Wait()
	//close(ch2)
	for ret := range b{		// 什么时候for range会退出？	b通道被关闭的时候
		fmt.Println(ret)
	}

}
