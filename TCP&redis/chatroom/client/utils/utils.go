package utils

import (
	"client/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte	//这是传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header errpr")
		//fmt.Printf("conn.Read() err=%v\n", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4]) //byte切片转换成uint32
	//根据pkgLen读取消息内容(从conn里面读pkgLen直接扔到buf里面)
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body errpr")
		//fmt.Printf("conn.Read(buf) err=%v\n", err)
		return
	}
	//把pkgLen反序列化成->message.Message
	//！！！技术就是一层窗户纸 &mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(buf) err=%v\n", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write(buf) err=%v\n", err)
		return
	}
	return
}