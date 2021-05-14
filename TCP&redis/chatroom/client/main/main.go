package main

import (
	"client/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续选择菜单
	var loop = true

	for loop {
		fmt.Println("--------------------欢迎登陆多人聊天系统--------------------")
		fmt.Println("\t\t\t 1、登陆聊天室")
		fmt.Println("\t\t\t 2、注册用户")
		fmt.Println("\t\t\t 3、退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\v", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1、创建一个UserProcess的实例
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
			 fmt.Printf("UserProcess login err=%v\n",err)
			 return
			}

			//loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			os.Exit(0)
			//fmt.Println("退出系统")
			//loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	//根据用户的输入，显示新的提示信息
	/*if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		//这里需要我们重新调用
		err := login(userId, userPwd)
		if err != nil {
			fmt.Printf("登录失败 err=%v\n",err)
		} else {
			//fmt.Println("登录成功")
		}

	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑...")
	}*/

}
