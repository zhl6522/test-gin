package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc");err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error":"err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		// 配置版本不一样 SetCookie传参不一致
		//c.SetCookie("abc", "123", 60, "/", "localhost", 0, false, true)
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(http.StatusOK, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data":"home"})
	})
	r.Run()
}
