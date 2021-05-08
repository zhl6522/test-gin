package main

import (
	"fmt"
	"io"
	"net" //做网络socket开发时，net包含有我们需要的所有方法和函数
	"time"
)

func process(conn net.Conn) {
	defer conn.Close() //如若不关闭 可能因为链接过多而没有释放，导致别的服务器链接不上来了

	for true {
		//创建一个切片
		buf := make([]byte, 1024)
		//conn.Read
		//1、等待客户端通过conn发送数据
		//2、如果客户端没有write[发送]，那么协程就会阻塞在这里
		//fmt.Printf("服务器在等待客户端 %s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("客户端 %s 已退出\n", conn.RemoteAddr().String())
			return
		}
		if err != nil {
			fmt.Printf("服务器的Read err=%v\n", err)
			return //!!! 协程里面报错了就需要return 不然一直报错
		}
		//3、显示客户端发送的内容到服务器的终端
		fmt.Printf("%v %v\n", conn.RemoteAddr().String(), time.Now().Format("2006-01-02 15:04:05"))
		fmt.Print(string(buf[:n])) //如若不写这个n，后面不知道的内容也会被读取进去(make())
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1、 tcp表示使用网络协议是tcp
	//2、0.0.0.0:8888 表示在本地监听8888端口（写0.0.0.0 IPV4、IPV6均可）
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen
	for true {
		//fmt.Println("循环等待客户端链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备起一个协程，为客户端服务
		go process(conn)
	}
	fmt.Printf("listen suc=%v\n", listen)

	//如若还没写客户端， cmd里请求 telnet 127.0.0.1 8888
}
