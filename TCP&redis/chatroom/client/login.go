package main

import (
	"client/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//fmt.Printf("你输入的用户id=%d 用户密码=%v\n", userId, userPwd)
	//return nil

	//1、链接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Printf("client err=%v\n", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2、准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3、创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4、将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
	 fmt.Printf("json.Marshal err=%v\n",err)
	 return
	}
	//5、把data赋给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n",err)
		return
	}
	//7、这个时候，data就是我们要发送的消息
	//7.1、为防止网络丢包，先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
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

	//fmt.Printf("客户端发送消息的长度=%d 内容=%v\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("conn.Write(data) err=%v\n",err)
		return
	}
	//time.Sleep(10*time.Second)
	//fmt.Println("休息10s")
	//这里还需要处理服务器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
	 fmt.Printf("readPkg(conn) err=%v\n",err)
	 return
	}
	//将mes的data部分反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
	 fmt.Printf("json.Unmarshal err=%v\n",err)
	 return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	
	return
}
