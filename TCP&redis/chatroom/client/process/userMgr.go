package process

import (
	"client/client/model"
	"client/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsersMap map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser //我们在用户登录成功后，完成对curUser初始化
//var UserId int

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsersMap
	fmt.Println("当前你在线用户列表：")
	for id, user := range onlineUsersMap {
		if user.UserStatus == 0 {
			fmt.Printf("用户id：\t%d\n", id)
		}
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsersMap[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onlineUsersMap[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
