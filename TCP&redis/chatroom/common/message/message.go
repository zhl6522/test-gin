package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	LogutMesType            = "LogutMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码 200表示登录成功
	UsersId []int  `json:"usersId"` //增加字段，保存用户id的切片
	Error   string `json:"error"`   //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//增加一个SmsMes发送消息
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体，继承
}

//客户端退出
type LogutMes struct {
	UserId int `json:"userId"`
}

type LogoutResMes struct {
	Code    int    `json:"code"`    //返回状态码 200表示登录成功
	UsersId []int  `json:"usersId"` //增加字段，保存用户id的切片
	Error   string `json:"error"`   //返回错误信息
}
