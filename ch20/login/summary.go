package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 1、json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"someJSON", "status":200})
	})
	// 2、结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct{
			Name	string
			Message	string
			Number	int
		}
		msg.Name="root"
		msg.Message="message"
		msg.Number=123
		c.JSON(http.StatusOK, msg)
	})
	// 3、XML	天气预报抓取转义
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message":"ok"})
		// 4、YAML响应	作为配置文件
		//c.YAML(http.StatusOK, gin.H{"name":"zhl"})
		// 5、protobuf格式，谷歌开发的高效存储读取工具	作为传输数据
		/*reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label:&label,
			Reps:reps,
		}
		c.ProtoBuf(http.StatusOK, data)*/
	})


	r.Run()

	// curl http://127.0.0.1:8080/root/admin
}
