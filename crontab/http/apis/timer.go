package apis

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"0924/http/client"
	"time"
)

//延时消费回调
//参数
/*
{
	"delay":"3600s"/"5m"/"1h"/"1d"最大支持天为单位 表示延迟执行的时间
	"request_url":"https://xxx.com/xxx" 用于执行回调的url
	"request_params":{"name":"rpc"} json格式参数列表或对象
}
*/
func HandleTimer(w http.ResponseWriter, r *http.Request){
	//获取参数
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
    	w.Write([]byte("接收参数出错了"))
        return
    }

    //json to map
    var params map[string]interface{}
    if err = json.Unmarshal(body, &params); err != nil {
        w.Write([]byte("参数格式有错误"))
        return
    }
   	log.Println(params)
    //存储任务

    //请求
    timeStr,_ := params["delay"].(string)
    requestUrl,_ := params["request_url"].(string)
    requestParams,_ := json.Marshal(params["request_params"])
	//携程
    go exec(timeStr,requestUrl,string(requestParams))
	w.Write([]byte("任务已加载..."))
}

//执行任务
func exec(timeStr, requestUrl, requestParams string){
	duration,err := time.ParseDuration(timeStr)
	if err != nil{
		log.Println(err)
	}
	timer := time.NewTimer(duration)
	log.Println(<-timer.C,"开始执行任务...")
	result,err := client.HttpPost(requestUrl,requestParams)
    if err != nil{
    	log.Println(err)
    }
    log.Println(result)
}
