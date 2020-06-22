package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端
func processConn(conn net.Conn) {
	defer conn.Close()
	for true {
		reader := bufio.NewReader(os.Stdin)
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failled,err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Println("请回复：")
		msg, _ := reader.ReadString('\n')		// 读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	// 1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:2000 failed,err:", err)
		return
	}
	// 2、等待别人来跟我建立连接
	for true {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		// 3、与客户端通信
		go processConn(conn)
	}
}
