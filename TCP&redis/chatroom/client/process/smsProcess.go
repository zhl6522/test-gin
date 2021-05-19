package process

import (
	"client/client/utils"
	"client/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1、创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//2、创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	//smsMes.User = curUser	//curUser的链接一页会赋给对方 所以不能这么写
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus
	//3、序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
	 fmt.Printf("SendGroupMes json.Marshal err=%v\n",err)
	 return
	}
	mes.Data = string(data)
	//4、对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
	 fmt.Printf("SendGroupMes json.Marshal err=%v\n",err)
	 return
	}
	//5、将序列化的mes发送给服务器
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
	 fmt.Printf("SendGroupMes WritePkg err=%v\n",err)
	 return
	}
	return
}