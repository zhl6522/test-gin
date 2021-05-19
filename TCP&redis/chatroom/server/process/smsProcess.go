package process

import (
	"client/common/message"
	"client/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

//转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的onlineUsers map[int]*UserProcess，将消息转发出去
	//取出mes的内容 SMSMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Printf("SendGroupMes Server json.Unmarshal err=%v\n", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("SendGroupMes Server json.Marshal err=%v\n", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//这里过滤一下自己，即不要发给自己
		if smsMes.UserId == id {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Printf("服务器端转发消息失败 err=%v\n", err)
	}
}
