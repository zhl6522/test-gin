package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client err=%v\n", err)
		return
	}
	for true {

		//功能一：客户端可以发送单行数据，然后就退出
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ReadString err=%v\n", err)
			return
		}

		//功能二：客户端可以发送多行数据，终端输入"exit"就退出
		line = strings.Trim(line, " \r\n")
		line = strings.Trim(line, " \n")
		if line == "exit" {
			fmt.Printf("客户端 %s 已退出\n", conn.RemoteAddr().String())
			break
		} else if line == "exit" { //windows下cmd的退出
			fmt.Printf("客户端 %s 已退出\n", conn.RemoteAddr().String())
			break
		}

		//再将line发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Printf("conn.Write err=%v\n", err)
			return
		}
		fmt.Printf("客户端发送了 %d 字节的数据，内容为：%s\n", n, line)
	}
}
