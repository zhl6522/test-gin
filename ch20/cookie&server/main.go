package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 客户端设置cookie
			// maxAge int, 单位为秒
			// path, cookie所在目录
			// domain string, 域名
			// secure 是否只能通过https访问
			// httpOnly bool 是否允许别通过js获取自己的cookie
			// 配置版本不一样 SetCookie传参不一致
			c.SetCookie("key_cookie", "val_cookie", 60, "/", "localhost", 0,false, true)
			//c.SetCookie("key_cookie", "val_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie的值是：%s\n", cookie)
	})
	r.Run()
}
