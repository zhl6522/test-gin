package process

import (
	"client/common/message"
	"client/server/model"
	"client/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//核心代码...
	//1、先从mes中取出Data，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	//2、声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//3、声明一个registerResMes，并完成赋值
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
		} else {
			registerResMes.Code = 506
		}
		registerResMes.Error = err.Error()
	} else {
		registerResMes.Code = 200
		fmt.Println("用户注册成功")
	}

	//4、将registerResMes序列化
	data, err := json.Marshal(registerResMes) //返回的data是切片
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、将data赋值给resMes
	resMes.Data = string(data)
	//6、对resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
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
	//3、声明一个loginResMes，并完成赋值
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
		} else {
			loginResMes.Code = 505
		}
		loginResMes.Error = err.Error()
	} else {
		loginResMes.Code = 200
		//用户登录成功，我们就把该用户放到userMgr中
		//将登录成功的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//将当前用户的id放到loginResMes.UsersId中
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user.UserName, "登录成功")
	}
	//4、将loginResMes序列化
	data, err := json.Marshal(loginResMes) //返回的data是切片
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}
	//5、将data赋值给resMes
	resMes.Data = string(data)
	//6、对resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
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
