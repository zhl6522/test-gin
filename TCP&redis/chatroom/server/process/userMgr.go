package process

import "fmt"

//因为UserMgr实例在服务器有且仅有一个并且在很多地方都会使用到，因此我们将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对UserMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加&修改
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//查询
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok {	//说明你要查找的这个用户当前不在线。
		err = fmt.Errorf("用户%d 当前不在线", userId)
		return
	}
	return
}
