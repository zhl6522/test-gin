package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)


var (g errgroup.Group)

func router01() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,
			"error":"Welcome 01"})
	})

	return r
}

func router02() http.Handler {
	r:=gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"code":http.StatusOK,
				"error":"Welcome server 02"})
	})

	return r
}

func main() {
	server01 := &http.Server{
		Addr:":8001",
		Handler: router01(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	server02 := &http.Server{
		Addr:":8002",
		Handler:router02(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	g.Go(func() error{
		return server01.ListenAndServe()
	})

	g.Go(func() error{
		return server02.ListenAndServe()
	})

	if err := g.Wait();err != nil {
		log.Fatal("",err)
	}
}