package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "zhl") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func Async(c *gin.Context) {
	// 创建要在goroutine中使用的副本
	cCp := c.Copy()
	go func() {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// 这里使用你创建的副本
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
func Sync(c *gin.Context) {
	// simulate a long task with time.Sleep(). 5 seconds
	time.Sleep(5 * time.Second)

	// 这里没有使用goroutine，所以不用使用副本
	log.Println("Done! in path " + c.Request.URL.Path)
}

func Redirect2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func Search(c *gin.Context) {
	//name := c.DefaultQuery("name", "zhl")
	//name := c.Query("name")
	//name := c.Param("name")	///user/search/:username/:address
	name := c.PostForm("name")
	address := c.PostForm("address")
	c.JSON(http.StatusOK, gin.H{
		"message":"ok",
		"name":name,
		"address":address,
	})
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Println("err:%v", err)
		return
	}
	name :="小王子"
	err =t.Execute(w, name)
	if err != nil {
		fmt.Println("file err:%v", err)
		return
	}
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello Gin")
}