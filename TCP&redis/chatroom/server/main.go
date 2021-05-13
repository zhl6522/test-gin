package main

import (
	"client/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
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

//serverProcessLogin函数：处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码...
	//1、先从mes中取出Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}

	//2、声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//3、声明一个loginMresMes，并完成赋值
	var loginMresMes message.LoginResMes

	if loginMes.UserId == 123 && loginMes.UserPwd == "qwert" {
		//合法
		loginMresMes.Code = 200
	} else {
		//不合法
		loginMresMes.Code = 500
		loginMresMes.Error = "该用户不存在，请先注册..."
	}
	//4、将loginMresMes序列化
	data, err := json.Marshal(loginMresMes)	//返回的data是切片
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、将data赋值给resMes
	resMes.Data = string(data)
	//6、对resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
	 fmt.Printf("json.Marshal err=%v\n",err)
	 return
	}
	//7、发送data，我们将其封装到writePkg函数中
	err = writePkg(conn, data)
	return
}

//serverProcessMes函数：根据客户端发送消息种类不同，决定调用相应的函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理和客户端的通讯
func process(conn net.Conn) {
	//！！！这里需要延时关闭
	defer conn.Close()

	//循环读客户端发送的信息
	for true {
		//这里我们将读取的数据包直接封装成一个函数readPkg()，返回Message，Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Printf("readPkg err=%v\n", err)
				return
			}
		}
		err = serverProcessMes(conn, &mes)

		fmt.Println("消息内容=", mes)
	}

}

func main() {

	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Printf("net.Listen err=%v\n", err)
		return
	}
	defer listen.Close() //延时关闭listen

	for true {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			//这里不退出 一个链接出错 其他的不退出
			fmt.Printf("listen.Accept() err=%v\n", err)
		}

		//链接成功后，则启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
