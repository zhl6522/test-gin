package process

import (
	"client/common/message"
	"client/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//serverProcessLogin函数：处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	//因为使用呢了分层模式(MVC)，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
