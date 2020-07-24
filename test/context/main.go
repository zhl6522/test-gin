package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context？
var wg sync.WaitGroup
//var notify bool
//var exitChan chan bool	//没有初始化
var exitChan = make(chan bool,1)

func f() {
	defer wg.Done()
	LOOP:
	for true {
		/*if notify {
			break
		}*/
		select {
		case <-exitChan:
			//break		// 跳出重复for
			break LOOP
		default:
		}
		fmt.Println("zhl")
		time.Sleep(time.Millisecond*500)
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second*5)
	// 如何通知子goroutine退出
	//notify = true
	exitChan<-true
	wg.Wait()
}