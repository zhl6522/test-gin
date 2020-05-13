package main

import (
	"fmt"
	"os"
)

// 学生管理系统

var smr studentMgr		// 声明一个全局的变量学生管理对象smr
func init() {
	fmt.Println("----------------welcome sms!-----------------")
}
// 菜单函数
func showMenu() {
	fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、编辑学生
		4、删除学生
		5、退出
		`)
}

func main() {
	smr = studentMgr{		// 修改全局的那个变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Print("请输入序号：")
		var choise int64
		fmt.Scanln(&choise)
		fmt.Println("你输入的是：", choise)
		switch choise {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.delStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("不要闹~")
		}
	}
}
