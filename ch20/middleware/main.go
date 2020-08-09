package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始i执行了")
		// 设置变量到context的key中，可以通过get()取
		c.Set("request", "中间件")
		// 执行中间件
		c.Next()	//最重要的步骤
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("程序执行时间：", since)
}

func main() {
	// 创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MyTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Use(MiddleWare())
	// {}为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request":req})
		})
		// 根路由后面是定义的局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request":req})
		})
	}
	r.Run()
}
func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 *time.Second)
}
func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}