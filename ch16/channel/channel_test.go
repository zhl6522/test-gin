package channel

import (
	"fmt"
	"os"
	"testing"
)

//引⽤类型 channel 是 CSP 模式的具体实现，⽤于多个 goroutine 通讯。其内部实现了
//同步，确保并发安全。
//默认为同步模式，需要发送和接收配对。否则会被阻塞，直到另⼀⽅准备好后被唤醒。
func TestChannel(t *testing.T) {
	data := make(chan int) // 数据交换队列
	exit := make(chan bool) // 退出通知
	go func() {
		for d := range data { // 从队列迭代接收数据，直到 close 。
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true // 发出退出通知。
	}()
	data <- 1 // 发送数据。
	data <- 2
	data <- 3
	close(data) // 关闭队列。
	fmt.Println("send over.")
	<-exit // 等待退出通知。
}

//异步⽅式通过判断缓冲区来决定是否阻塞。如果缓冲区已满，发送被阻塞；缓冲区为空，
//接收被阻塞。
//通常情况下，异步 channel 可减少排队阻塞，具备更⾼的效率。但应该考虑使⽤指针规
//避⼤对象拷⻉，将多个元素打包，减⼩缓冲区⼤⼩等。
func TestChannel2(t *testing.T) {
	data := make(chan int, 3) // 缓冲区可以存储 3 个元素
	exit := make(chan bool)
	data <- 1 // 在缓冲区未满前，不会阻塞。
	data <- 2
	data <- 3
	go func() {
		for d := range data { // 在缓冲区未空前，不会阻塞。
			fmt.Println(d)
		}
		exit <- true
	}()
	data <- 4 // 如果缓冲区已满，阻塞。
	data <- 5
	close(data)
	<-exit
}

//如果需要同时处理多个 channel，可使⽤ select 语句。它随机选择⼀个可⽤ channel 做
//收发操作，或执⾏ default case。
//在循环中使⽤ select default case 需要⼩⼼，避免形成洪⽔。
func TestChannel3(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			select { // 随机选择可⽤ channel，接收数据。
			case v, ok = <-a: s = "a"
			case v, ok = <-b: s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		select { // 随机选择可⽤ channel，发送数据。
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	select {} // 没有可⽤ channel，阻塞 main goroutine。
	/*
	b 3
	a 0
	a 1
	a 2
	b 4
	*/
}