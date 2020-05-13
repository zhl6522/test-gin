package main

import (
	"fmt"
	"os"
)

/*
	函数版学生系统
	写一个系统能够新增/查看、编辑学生
*/

type student struct {
	id		int
	name	string
	age		int
	score	int
}

// newStatudent 是student类型的构造函数
func newStatudent(id int, name string, age int, score int) *student {
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

var (
	allStudent map[int]*student		// 变量声明
)

func showAllStudent() {
	// 把所有学生都打印出来
	for k,v := range allStudent {
		fmt.Printf("学生的学号：%d，姓名：%v，年龄：%d，分数：%d\n", k, v.name, v.age, v.score)
	}
}

func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1、创建一个新的学生
	var (
		stuId		int
		stuName		string
		stuAge		int
		stuScore	int
	)
	// 1.1、获取用户输入
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&stuId)
	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&stuName)
	fmt.Print("请输入学生的年龄:")
	fmt.Scanln(&stuAge)
	fmt.Print("请输入学生的分数:")
	fmt.Scanln(&stuScore)
	// 1.2、造学生（调用student的构造函数）
	newStu := newStatudent(stuId, stuName, stuAge, stuScore)
	// 追加到allStudent这个map中
	allStudent[stuId] = newStu
}

func editStudent() {
	var (
		stuId		int
		stuName		string
		stuAge		int
		stuScore	int
	)
	// 1.1、获取用户输入
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&stuId)
	fmt.Print("请输入学生的姓名:")
	fmt.Scanln(&stuName)
	fmt.Print("请输入学生的年龄:")
	fmt.Scanln(&stuAge)
	fmt.Print("请输入学生的分数:")
	fmt.Scanln(&stuScore)
	editStu := &student{
		id:    stuId,
		name:  stuName,
		age:   stuAge,
		score: stuScore,
	}
	allStudent[stuId] = editStu
}


func delStudent() {
	// 1、请用户输入要删除的学号
	var (
		delId int
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&delId)
	// 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, delId)
}
func init() {
	fmt.Println("欢迎光临学生管理系统！")
}

func main() {
	allStudent = make(map[int]*student, 50)		// 初始化（开辟内存空间）
	for {
		// 1、打印菜单
		fmt.Println(`
			1、查看所有学生
			2、新增学生
			3、编辑学生
			4、删除学生
			5、退出
		`)
		fmt.Print("请输入你要干啥：")
		// 2、等待用户选择要做什么
		var choise int
		fmt.Scanln(&choise)
		fmt.Printf("您选择了%d这个选项！\n", choise)
		// 3、执行对应的函数
		switch choise {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			editStudent()
		case 4:
			delStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("不要闹~")

		}
	}

}
