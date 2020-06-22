package main

import (
	"fmt"
	"net"
	"strings"
)

// UDP server

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:net.IPv4(127,0,0,1),
		Port:4000,
	})
	if err != nil {
		fmt.Println("listen udp failed,err:", err)
		return
	}
	defer conn.Close()
	// 不需要建立连接，直接收发数据
	var data [1024]byte		//数组 data[:]切片
	for true {
		n,addr,err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed,err:", err)
			return
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		reply := strings.ToUpper(string(data[:n]))
		// 发送数据
		_, err = conn.WriteToUDP([]byte(reply), addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
