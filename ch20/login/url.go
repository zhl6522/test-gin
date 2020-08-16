package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必填字段
	User 		string	`form:"username"json:"user"uri:"user"xml:"user"binding:"required"`
	Password	string	`form:"password"json:"password"uri:"password"xml:"password"binding:"required"`
}

func main() {
	r := gin.Default()
	// JSON绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自行推断
		if err := c.ShouldBindUri(&login);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:":err.Error()})
			return
		}
		// 判断用户密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status:":"304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status:":"200"})
	})

	r.Run()

	// curl http://127.0.0.1:8080/root/admin
}
