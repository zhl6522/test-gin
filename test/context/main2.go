package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context？
var wg sync.WaitGroup
//var notify bool
//var exitChan chan bool	//没有初始化
//var exitChan = make(chan bool,1)

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for true {
		/*if notify {
			break
		}*/
		select {
		case <-ctx.Done():
			//break		// 跳出重复for
			break LOOP
		default:
		}
		fmt.Println("mumu2333")
		time.Sleep(time.Millisecond*500)
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for true {
		/*if notify {
			break
		}*/
		select {
		case <-ctx.Done():
			//break		// 跳出重复for
			break LOOP
		default:
		}
		fmt.Println("zhl")
		time.Sleep(time.Millisecond*500)
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()		// 当我们取完需要的整数后调用cancel


	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*5)
	// 如何通知子goroutine退出
	cancel()
	wg.Wait()
}