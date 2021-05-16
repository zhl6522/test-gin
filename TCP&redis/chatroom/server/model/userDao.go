package model

import (
	"client/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后，就初始化一个UserDao实例，把它做成全局的变量，在需要redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体，完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考一下在UserDao应该提供哪些方法给我们
//1、根据影虎id，返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	//通过给定userId去redis查询这个用户
	res, err := redis.String(conn.Do("Hget", "users", userId))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应userId
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里我们需要把res反序列化成User实例
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func (this UserDao) Register(user *message.User) (err error) {
	//先从UserDao的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时说明userId在redis还没有，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	//入库
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Printf("保存注册用户失败， err=%v\n", err)
		return err
	}
	return
}

//完成登录的校验 Login
func (this UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
