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

//这里我们通知所有在线用户的方法
//userId通知其他在线用户，他上线了
func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	//遍历UserMgr.onlineUsers，然后一个一个发送NotifyUserStatusMes消息
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int)  {
	//组装NotifyUserStatusMes消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
	 fmt.Printf("json.Marshal err=%v\n",err)
	 return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
	 fmt.Printf("json.Marshal err=%v\n",err)
	 return
	}

	//发送，创建Transfer实例
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
	 fmt.Printf("NotifyMeOnline err=%v\n",err)
	 return
	}

}

func (this *UserProcess) ServerProcessLogout(mes *message.Message) {
	//客户端退出
	var logoutMes message.LogutMes
	err := json.Unmarshal([]byte(mes.Data), &logoutMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n",err)
		return
	}
	userMgr.DelOnlineUser(logoutMes.UserId)

	this.NotifyOtherOfflineUser(logoutMes.UserId)
	fmt.Println(logoutMes.UserId, "退出成功")
}

func (this *UserProcess) NotifyOtherOfflineUser(userId int) {
	//遍历UserMgr.onlineUsers，然后一个一个发送NotifyUserStatusMes消息
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOffline(userId)
	}
}

func (this *UserProcess) NotifyMeOffline(userId int)  {
	//组装NotifyUserStatusMes消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOffline
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n",err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n",err)
		return
	}

	//发送，创建Transfer实例
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("NotifyMeOnline err=%v\n",err)
		return
	}
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
	//因为使用了分层模式(MVC)，我们先创建一个Transfer实例，然后读取
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
		//通知其他在线用户，他上线了
		this.NotifyOtherOnlineUser(loginMes.UserId)
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
	//因为使用了分层模式(MVC)，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
