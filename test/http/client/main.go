package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http client端
// 共用一个client适用于 青丘族比较频繁
var (			// 如果设置为5分钟拉取一次数据，需要改成false 复用连接，不然大概一个月后，所有的长连接都没有被释放，就拉取不到数据了。
	client = http.Client{
		Transport:&http.Transport{
			DisableKeepAlives:false,
		},
	}
)

func main() {
	//get()
	post()
}

func post() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=笨笨&age=28"
	// json
	contentType := "application/json"
	data := `{"name":"笨笨","age":28}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
	 fmt.Printf("post failed, err:%v\n",err)
	 return
	}
	defer resp.Body.Close()
	body, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed, err:%v\n",err)
		return
	}
	fmt.Println(string(body))
}

func get() {
	/*resp, err := http.Get("http://127.0.0.1:9090/query/?name=zhl&age=25")
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}*/
	// URL param
	data := url.Values{}		// url values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/query/")
	data.Set("name", "沐沐")
	data.Set("age", "28")
	urlObj.RawQuery = data.Encode()		// URL encode之后的URL
	fmt.Println(urlObj.RawQuery)
	fmt.Println(urlObj.String())
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	/*resp, err := http.DefaultClient.Do(req)
	if err != nil {
	 fmt.Printf("get url failed, err:%v\n",err)
	 return
	}*/
	// 请求不是特别频繁，用完就关闭该连接
	// 禁用KeepAlive的client（禁用长连接）	每次都会新建一个连接去发起请求
	/*tr := &http.Transport{
		DisableKeepAlives:true,
	}
	client := http.Client{
		Transport:tr,
	}*/
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed, err:%v\n",err)
		return
	}

	defer resp.Body.Close()		// 关闭ressp.Body连接
	// 从resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	body, err :=ioutil.ReadAll(resp.Body)		// 客户端读出服务端返回的响应的body
	if err != nil {
	 fmt.Printf("read resp.body failed, err:%v\n",err)
	 return
	}
	fmt.Println(string(body))
}
