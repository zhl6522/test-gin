package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client端

func main() {
	// 1、与server端建立连接
	conn,err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:2000 failed，err:", err)
		return
	}
	// 2、发送数据
	//var msg string
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("请说话：")
		msg, _ := reader.ReadString('\r')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	/*if len(os.Args) < 2 {
		msg = "Hello world!"
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))*/
	/*if err != nil {
		fmt.Println("Write failed, err:", err)
		return
	}*/
	//fmt.Println(n)
	conn.Close()
}
