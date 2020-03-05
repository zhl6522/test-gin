package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-gin/controller"
)

type User struct {
	Name string
	Gender string
	Age int
}

func main() {
	// 新建一个没有任何默认中间件的路由
	//r := gin.New()
	r := gin.Default()
	r.LoadHTMLGlob("./templates/**/*")

	//也可以这么写
	/*
	shopGroup := r.Group("/shop")
	shopGroup.Use(controller.StatCost())
	{}
	*/
	shopGroup := r.Group("/shop", controller.StatCost())
	{
		shopGroup.GET("/test", controller.Redirect2)
	}

	r.GET("/test3", controller.StatCost(), func(c *gin.Context) {	//给/test3路由单独注册中间件（可注册多个）
		name := c.MustGet("name").(string)		//从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	// 注册一个全局中间件
	r.Use(controller.StatCost())
	r.GET("/test1", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", controller.Redirect2)
	r.GET("file", controller.File)
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", controller.UploadFile)
	r.GET("files", controller.Files)
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	//r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/uploads", controller.UploadFiles)
	r.POST("/loginJSON", controller.LoginUser)

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name string
			Gender string
			Age int
		}
		msg.Name = "benben"
		msg.Gender = "woman"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.POST("/user/search", controller.Search)

	r.Static("/static", "./static")
	//===========================
	http.HandleFunc("/hello", controller.SayHello)
	//r.LoadHTMLGlob("./templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	u1 := User{
		Name:"zhl",
		Gender:"man",
		Age:19,
	}
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title":"hello GoLand!",
			"user":u1.Name,
			"gender":u1.Gender,
			"age":u1.Age,
		})
	})
	m1 :=make(map[string]interface{})
	m1["name"] = "mumu"
	m1["gender"] = "man"
	m1["age"] = "25"
	//m1 := map[string]interface{}{
	//	"name":"zhl",
	//	"gender":"man",
	//	"age":19,
	//}
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "hello GoLand!",
			"user":m1["name"],
			"gender":m1["gender"],
			"age":m1["age"],
		})
	})
	r.Run()
}