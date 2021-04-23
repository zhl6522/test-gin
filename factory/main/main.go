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
	fmt.Println(*stu2, stu2, stu2.Name, stu2.GetScore()) //不能直接stu2.score调取

	//继承
	var st = Sta{}
	st.Name = "yoyo" //原本是这样的：st.student.Name="yoyo"
	st.age = 22
	st.Info()
	//结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大小或者小写的字段、方法都可以使用。
	st.studying()

	var st2 = Sta2{}
	//当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
	st2.Name = "mumu"
	st2.student.Name = "mumu~" //st2.student.Name未定义的话，默认为空字符串
	st2.age = 20
	st2.int = 5 //匿名字段直接调用
	st2.Info()
	st2.studying()

	tv := TV{Goods{"电视机001", 2688.8}, Brand{"海尔", "山东"}}
	fmt.Printf("TV=%v tv商品名=%v\n", tv, tv.Goods.Name)
	//为了保证代码的简洁性，建议大家尽量不适用多重继承

}

//继承
type student struct {
	Name string
	age  int
}

func (st *student) Info() {
	fmt.Printf("学生名：%v 年龄：%d\n", st.Name, st.age)
}

type Sta struct {
	student
}

func (stu *Sta) studying() {
	fmt.Printf("%v学习中...\n", stu.Name)
}

type Sta2 struct {
	student
	Name string
	int  //匿名字段也是基本数据类型，一个结构体最毒多只有一个匿名字段，如果需要多个int字段，则必须给int字段指定名称。
}

func (stu *Sta2) studying() {
	fmt.Printf("%v学习中...\n", stu.Name)
}

//多重继承
type Goods struct {
	Name  string
	Place float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}
