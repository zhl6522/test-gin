package main

import (
	"fmt"
	"net"
)

// socket_stick/client/main.go
// https://www.liwenzhou.com/posts/Go/15_socket/
// 解决"粘包"的办法：使用大端小端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
		//time.Sleep(time.Millisecond*50)
	}
}
