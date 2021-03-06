package main

import (
	"flag"
	"fmt"
)

func main() {
	//os.Args使用 如果命令行传参位置换了就无法获取到数据
	//flag包可以根据指定参数获取对应的值
	var user, pwd, host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" 就是 -u 指定参数
	//"" 默认值
	//"用户名，默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "127.0.0.1", "主机名，默认为127.0.0.1")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//转换，必须调用该方法
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
