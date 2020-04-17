package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-gin/controller"
	//"golang.org/x/autotls"
	//"golang.org/x/crypto/acme/autocert"
)

type User struct {
	Name string
	Gender string
	Age int
}

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// 新建一个没有任何默认中间件的路由
	//r := gin.New()
	r := gin.Default()



	/*//支持Let's Encrypt证书
	r.GET("/hello", controller.Hello)
	//log.Fatal(autotls.Run(r, "ahfuzl.com", "zhl.com"))
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("E:/www/go_project/src/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))*/

	//中间件中使用Goroutines
	r.GET("/long_async", controller.Async)
	r.GET("/long_sync", controller.Sync)

	//使用BasicAuth()（验证）中间件
	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.LoadHTMLGlob("./templates/**/*")

	//返回第三方获取的数据 文件会下载
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			fmt.Println("err:%s", err.Error)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; path="E:\www\go_project\src\test-gin"; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

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
	//r.StaticFile("/index.css", "./static/index.css")
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