package main

import (
	"client/server/model"
	"fmt"
	"net"
	"time"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//！！！这里需要延时关闭
	defer conn.Close()

	//这里调用总控
	processor := &Processor{Conn: conn}
	err := processor.Process2()
	if err != nil {
		fmt.Printf("客户端和服务器端通信协程错误=%v\n", err)
		return
	}

}

func init() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool("localhost:6379", 10, 0, 300*time.Second)
	initUserDao()
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool本身就是一个全局的变量
	model.MyUserDao = model.NewUserDao(pool)
}

/*
Administrator@MS-20170306-zhl MINGW64 /e/www/go_project/src/test-gin/TCP&redis/chatroom (master)
$ go build -o server.exe server/main/main.go server/main/processor.go server/main/redis.go
Administrator@MS-20170306-zhl MINGW64 /e/www/go_project/src/test-gin/TCP&redis/chatroom (master)
$ server.exe
*/
func main() {

	fmt.Println("服务器[新的结构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Printf("net.Listen err=%v\n", err)
		return
	}
	defer listen.Close() //延时关闭listen

	for true {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			//这里不退出 一个链接出错 其他的不退出
			fmt.Printf("listen.Accept() err=%v\n", err)
		}

		//链接成功后，则启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
