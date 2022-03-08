package main

import (
	"context"
	"github.com/ffhelicopter/tmm/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {




	router := gin.Default()

	//静态资源加载，本例为css,js文件以及资源图片
	router.StaticFS("/public", http.Dir("E:/www/go_project/pkg/mod/github.com/ffhelicopter/tmm/website/static"))
	router.StaticFile("favicon.ico", "./resources/favicon.ico")

	//导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("E:/www/go_project/pkg/mod/github.com/ffhelicopter/tmm/website/tpl/*/*")

	//reboot分组
	v := router.Group("/")
	{
		v.GET("/index.html", handler.IndexHandler)
		v.GET("/add.html", handler.AddHandler)
		v.POST("/postme.html", handler.PostmeHandler)
	}
	v2 := router.Group("/login")
	{
		v2.GET("/index.html", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"meg": "login->index.html",
			})
		})
		v2.GET("/add.html", handler.AddHandler)
		//v2.GET("/ss", tollbooth.LimitFuncHandler(lmt, HelloHandler))
	}
	//router.Run(":8080")
	//下面代码(Go1.8+版本支持)是为了优雅处理重启。
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	go func() {
		//监听请求
		if err :=srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	//优雅shutdown(或重启)服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)	//syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err :=srv.Shutdown(ctx);err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
		case<-ctx.Done():
	}
	log.Println("Server exiting")

}
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}