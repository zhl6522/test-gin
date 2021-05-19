package main

import (
	"client/common/message"
	process2 "client/server/process"
	"client/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//serverProcessMes函数：根据客户端发送消息种类不同，决定调用相应的函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	//看看是否能接收到客户端发送的群聊消息
	fmt.Println("mes=", mes)

	switch mes.Type {
		case message.LoginMesType:
			//处理登录逻辑
			//创建一个UserProcess实例
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册逻辑
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			//创建一个SmsProcess实例完成转发群聊消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		case message.LogutMesType:
			//客户端退出
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			up.ServerProcessLogout(mes)
		default:
			fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error) {
	//循环读客户端发送的信息
	for true {
		//这里我们将读取的数据包直接封装成一个函数readPkg()，返回Message，Err
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return err
			} else {
				fmt.Printf("readPkg err=%v\n", err)
				return err
			}
		}
		err = this.ServerProcessMes(&mes)
		if err != nil {
		 return err
		}
		fmt.Println("消息内容=", mes)
	}
	return
}
