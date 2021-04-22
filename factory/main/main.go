package main

import (
	"fmt"
	"test-gin/factory/model"
)

// 工厂模式
func main() {
	var stu = model.Student{
		Name:  "zhl",
		Score: 88,
	}
	fmt.Println(stu)
	stu2 := model.Newstudent("mumu", 88.8)
	fmt.Println(*stu2, stu2, stu2.Name, stu2.GetScore())	//不能直接stu2.score调取

	var st = Sta{}
	st.student.Name="yoyo"
	st.student.age=22
	st.Info()
	st.studying()

	var st2 = Sta2{}
	st2.student.Name="yoyo"
	st2.student.age=22
	st.Info()
	st2.studying()

}

//继承
type student struct {
	Name	string
	age		int
}

func (st *student)Info()  {
	fmt.Printf("学生名：%v 年龄：%d\n", st.Name, st.age)
}

type Sta struct {
	student
}

func (stu *Sta) studying() {
	fmt.Println("sta学习中...")
}

type Sta2 struct {
	student
}

func (stu *Sta2) studying() {
	fmt.Println("sta2学习中...")
}