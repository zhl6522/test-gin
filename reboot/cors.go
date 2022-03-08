package main

import "github.com/didip/tollbooth"

func main() {

	//中间件的使用
	//rate-limit 限流中间件
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMessage("服务器繁忙，请稍后再试...")
}
