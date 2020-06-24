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

func main() {
	http.HandleFunc("/hi/", f1)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
