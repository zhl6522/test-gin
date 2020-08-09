package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
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
