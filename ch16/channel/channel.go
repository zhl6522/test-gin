package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//用简单工厂模式打包并发任务和 channel。
func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()
	return c
}
func main() {
	t := NewTest()
	println(<-t) // 等待 goroutine 结束返回。

	//用 channel 实现信号量 (semaphore)。
	wg := sync.WaitGroup{}
	wg.Add(3)
	sem := make(chan int, 1)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sem <- 1 // 向 sem 发送数据，阻塞或者成功。
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
			<-sem // 接收数据，使得其他阻塞 goroutine 可以发送数据。
		}(i)
	}
	wg.Wait()

	//用 closed channel 发出退出通知。
	var wg2 sync.WaitGroup
	quit := make(chan bool)
	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			task := func() {
				println(id, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}
			for {
				select {
				case <-quit: // closed channel 不会阻塞，因此可用作退出通知。
					return
				default: // 执行正常任务。
					task()
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 5) // 让测试 goroutine 运行一会。
	close(quit) // 发出退出通知。
	wg2.Wait()
}
