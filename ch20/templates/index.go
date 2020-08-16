package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// html渲染
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("test-gin/ch20/**/*")
	//r.LoadHTMLFiles("test-gin/ch20/templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终json将title替换
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title":"我的标题"})
	})
}
