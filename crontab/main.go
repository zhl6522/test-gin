package main

import(
	"0924/http/server"
	"0924/http/utils"
	"0924/http/apis"
)

func main(){
	addr := utils.GetAddrFromFlag()
	server := server.NewServerMux()

	server.POST("/timer",apis.HandleTimer) //注册定时消费路由

	server.Run(addr)
}

