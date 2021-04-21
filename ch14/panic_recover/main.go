package main

import (
	"errors"
	"fmt"
)

func PanicVxExit() {
	defer func() {
		if err := recover(); err != nil {	//不推荐使用：1、形成僵尸服务进程 2、错误信息被忽略
			//错误恢复
			fmt.Println("recovered panic ", err)
			fmt.Println("发送邮件给ahfuzl@126.com")
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))	//从输出你可以看出defer处理了panic中抛出的error
}

func main() {
	PanicVxExit()
	fmt.Println("main() aaaaaaaabc")
}