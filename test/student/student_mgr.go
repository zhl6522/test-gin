package main

import (
	"bufio"
	"fmt"
	"os"
)

// 学生管理系统
// 有一个物件
// 1、它保存的一些数据	--> 结构体的字段
// 2、它有4个功能  		--> 结构体的方法

type student struct {
	id		int64
	name	string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}
// 查看学生
func (s studentMgr) showStudents()  {
	// 从s.allStudent这个map中把所有学生逐个拎出来
	for _,stu := range s.allStudent {	// stu是具体每一个学生
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}
// 增加学生
func (s studentMgr) addStudent() {
	// 1、根据用户输入的内容创建一个新的学生
	var (
		stuId		int64
		stuName		string
	)
	// 获取用户输入
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&stuId)
	// 查询该学号对应的学生信息，如果有提示已存在
	_,ok := s.allStudent[stuId];
	if ok {
		fmt.Println("该学号已被人使用")
		return
	}
	//fmt.Scanln(&stuName)	// 如果姓名里有空格，就没法录入全部
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入学生的姓名:")
	stuName,_ = reader.ReadString('\n')
	// 根据用户输入创建结构体对象
	newStu := student{
		id:stuId,
		name:stuName,
	}
	// 2、把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功！")
}
// 修改学生
func (s studentMgr) editStudent() {
	// 1、获取用户输入的学号
	var stuId		int64
	// 获取用户输入
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&stuId)
	// 2、展示该学号对应的学生信息，如果没有提示查无此人
	stuObj,ok := s.allStudent[stuId];
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	// 3、请输入修改后的学生名
	var newName		string
	//fmt.Scanln(&newName)	// 如果姓名里有空格，就没法录入全部
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入学生的新姓名:")
	newName,_ = reader.ReadString('\n')
	// 4、更新学生的姓名
	stuObj.name = newName	// 这时只是修改了stuObj对象的name值
	s.allStudent[stuId] = stuObj	// 更新map中的学生
}
// 删除学生
func (s studentMgr) delStudent() {
	// 1、请用户输入要删除的学生id
	var stuId		int64
	// 获取用户输入
	fmt.Print("请输入要删除学生的学号:")
	fmt.Scanln(&stuId)
	// 2、去map中查找该id， 如果没有打印查无此人
	_,ok := s.allStudent[stuId];
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3、有的话就确认删除
	fmt.Print("请输入Y/N:")
	var stuY		string
	fmt.Scanln(&stuY)
	// 4、确认就删除
	if stuY != "Y" && stuY != "y" {
		fmt.Println("你已取消删除.")
		return
	}
	delete(s.allStudent, stuId)
	fmt.Println("删除成功！")
}

