package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/test-gin/helloGin/controller"
	"net/http"
)

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func main() {
	// Engin
	router := gin.New()

	// 加载html文件，即template包下所有文件
	//router.LoadHTMLGlob("template/*")
	// 路由组
	user := router.Group("/user")
	{
		user.GET("/get/:id/:name", controller.GetUser)
		user.GET("/query", controller.QueryParam)
		user.POST("/insert", controller.InsertUser)
		user.POST("/update", controller.UpdateUser)
		user.GET("/view", controller.RenderView)
		user.GET("/delete/:id", controller.DeleteUser)
	}
	// Engin
	/*router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login/user/:name", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("read", readEndpoint)
	}*/

	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run("127.0.0.1:15001")
}

/*func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"id":id,
		"user":user,
	})
}*/

func hello(c *gin.Context) {
	println(">>>> hello gin start <<<<")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
