package main

import (
	"client/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header errpr")
		//fmt.Printf("conn.Read() err=%v\n", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4]) //byte切片转换成uint32
	//根据pkgLen读取消息内容(从conn里面读pkgLen个志杰扔到buf里面)
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		//err = errors.New("read pkg body errpr")
		//fmt.Printf("conn.Read(buf) err=%v\n", err)
		return
	}
	//把pkgLen反序列化成->message.Message
	//！！！技术就是一层窗户纸 &mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(buf) err=%v\n",err)
		return
	}
	//发送data本身
	n, err = conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Printf("conn.Write(buf) err=%v\n",err)
		return
	}
	return
}
