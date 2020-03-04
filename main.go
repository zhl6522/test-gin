package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func main() {
	r := gin.Default()

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
	r.POST("/user/search", func(c *gin.Context) {
		//name := c.DefaultQuery("name", "zhl")
		//name := c.Query("name")
		name := c.PostForm("name")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"message":"ok",
			"name":name,
			"address":address,
		})
	})

	r.Static("/static", "./static")
	//===========================
	http.HandleFunc("/hello", sayHello)
	r.LoadHTMLGlob("./templates/**/*")
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

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Println("err:%v", err)
		return
	}
	name :="小王子"
	err =t.Execute(w, name)
	if err != nil {
		fmt.Println("file err:%v", err)
		return
	}
}