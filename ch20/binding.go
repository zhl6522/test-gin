package main

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必填字段
	User 		string	`form:"username"json:"user"uri:"user"xml:"user"binding:"required"`
	Password	string	`form:"password"json:"password"uri:"password"xml:"password"binding:"required"`
}

func main() {
	r := gin.Default()
	// JSON绑定
	r.POST("/loginJSON", func(c *gin.Context) {
		// 生命接收的变量
		var json Login
		// 讲require的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json);err != nil {
			// 返回错误信息
		}
	})

	r.Run()
}
