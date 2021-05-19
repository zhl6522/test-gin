package process

import (
	"client/client/utils"
	"client/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面...
func showMenu(userId int) {
	user := fmt.Sprintf("-------------恭喜%d登录成功-------------", userId)
	fmt.Println(user)
	fmt.Println("-------------1、显示在线用户列表-------------")
	fmt.Println("-------------2、发送消息-------------")
	fmt.Println("-------------3、信息列表-------------")
	fmt.Println("-------------4、退出系统-------------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	//因为我们总会使用到SmsProcess实例，因此我们将其定义在switch外部
	smsProcess := &SmsProcess{}
	userProcess := &UserProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你要怼大家说的内容：")
		fmt.Scanf("%v\n", &content)
		err := smsProcess.SendGroupMes(content)
		if err != nil {
		 fmt.Printf("smsProcess.SendGroupMes(content) err=%v\n",err)
		 return
		}

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		userProcess.Logout(userId)

		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确...")
	}
}

//与服务器端保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for true {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("tf.ReadPkg err=%v\n", err)
			return
		}
		//如果读取到消息，又是下一步逻辑
		switch mes.Type {
			case message.NotifyUserStatusMesType:
				//有人上线了
				//1、取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				if err != nil {
					fmt.Printf("json.Unmmarshal err=%v\n",err)
					return
				}
				//2、把这个用户的信息状态保存到客户端map[int]User中
				updateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType:
				//接收群聊消息
				outputGroupMes(&mes)
		default:
				fmt.Println("服务器端返回一个未知的消息类型")
		}

	}
}
