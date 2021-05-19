package process

import (
	"client/common/message"
	"encoding/json"
	"fmt"
	"time"
)

func outputGroupMes(mes *message.Message) {
	//1、反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("%d %s\n%s", smsMes.UserId, time.Now().Format("2006-01-02 15:04:05"), smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
