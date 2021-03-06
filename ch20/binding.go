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
	r.POST("/loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将require的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json);err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error:":err.Error()})
			return
		}
		// 判断用户密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status:":"304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status:":"200"})

		//  curl http://127.0.0.1:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\", \"password\":\"admin\"}" -X POST
		//  curl http://127.0.0.1:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\", \"password\":\"root\"}" -X POST
	})

	r.Run()
}
