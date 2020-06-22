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
	defer conn.Close()
	// 2、发送数据
	//var msg string
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("请说话：")
		msg, _ := reader.ReadString('\n')		// 读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
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
