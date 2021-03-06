package process

import (
	"client/client/utils"
	"client/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

//客户端退出
func (this *UserProcess) Logout(userId int) {
	var mes message.Message
	mes.Type = message.LogutMesType
	var logoutmes message.LogutMes
	logoutmes.UserId = userId
	data, err := json.Marshal(logoutmes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、把data赋给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	//发送数据给服务端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("退出发送信息错误：conn.WritePkg(data) err=%v\n", err)
		return
	}
}

func (this *UserProcess) Register(userId int, userPwd, userName string) (err error) {
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
	mes.Type = message.RegisterMesType
	//3、创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//4、将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、把data赋给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送数据给服务端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("注册发送信息错误：conn.WritePkg(data) err=%v\n", err)
		return
	}
	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Printf("readPkg() err=%v\n", err)
		return
	}
	//将mes的data部分反序列化成loginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

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
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、把data赋给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//7、这个时候，data就是我们要发送的消息
	//7.1、为防止网络丢包，先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //pkgLen赋给buf切片
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(buf) err=%v\n", err)
		return
	}

	fmt.Printf("客户端发送消息的长度=%d 内容=%v\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("conn.Write(data) err=%v\n", err)
		return
	}
	//time.Sleep(10*time.Second)
	//fmt.Println("休息10s")
	//这里还需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Printf("readPkg() err=%v\n", err)
		return
	}
	//将mes的data部分反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	if loginResMes.Code == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")
		//显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, value := range loginResMes.UsersId {
			//如果要求不显示自己在线
			if value == userId {
				continue
			}
			fmt.Println("用户id:\t", value)
			//完成客户端onlineUsersMap的初始化
			user := &message.User{
				UserId: value,
				//UserName:   "",
				UserStatus: message.UserOnline,
			}
			onlineUsersMap[value] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程，该协程保持和服务器端的通讯，如果服务器有数据推送给客户端，则接收并显示在客户端的终端。
		go serverProcessMes(conn)

		//1、显示我们登录成功的菜单...[这里循环也可以外面client/main.go循环]
		for true {
			showMenu(userId)
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
