package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server端

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	str := "Hello 昌平！"
	w.Write([]byte(str))
}

func get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 对于GET请求，参数都放在URL上（query param），请求体中是没有数据的
	queryParam := r.URL.Query()		// 自动识别URL中的 query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))		// 服务端打印客户端发来的请求的body
	w.Write([]byte("ok"))
}

func post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	/*r.ParseForm()
	fmt.Println(r.PostForm)		// 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))*/
	// 2. 请求类型是application/json时从r.Body读取数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	 fmt.Printf("read request.Body failed, err:%v\n",err)
	 return
	}
	fmt.Println(string(body))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/hi/", f1)
	http.HandleFunc("/query/", get)
	http.HandleFunc("/post", post)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
