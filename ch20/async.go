package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

func main() {
	rwMutex()
	
	r := gin.Default()
	// 异步
	r.GET("long_async", func(c *gin.Context) {
		// 异步有时候会失效，需要用一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行："+copyContext.Request.URL.Path)
		}()
	})
	// 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行："+c.Request.URL.Path)
	})
	r.Run()
}

func rwMutex() {
	var rwm sync.RWMutex

	for i := 0; i <= 2; i++ {
		go func(i int) {
			fmt.Printf("go(%d) start lock\n", i)
			rwm.RLock()
			fmt.Printf("go(%d) locked\n", i)
			time.Sleep(time.Second * 2)
			rwm.RUnlock()
			fmt.Printf("go(%d) unlock\n", i)
		}(i)
	}
	// 先sleep一小会，保证for的goroutine都会执行
	time.Sleep(time.Microsecond * 100)
	fmt.Println("main start lock")
	// 当子进程都执行时，且子进程所有的资源都已经Unlock了
	// 父进程才会执行
	rwm.Lock()
	fmt.Println("main locked")
	time.Sleep(time.Second)
	rwm.Unlock()
}