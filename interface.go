package main

import "fmt"


//声明/定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Computer struct {}
type Phone struct {}
func (p Phone) Start() {
	fmt.Println("手机在工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机已停止")
}
type Camera struct {}
func (c Camera) Start() {
	fmt.Println("相机在工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机已停止")
}
//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明的所有方法）
func (c Computer) Working(usb Usb)  {	//usb变量既可以接收手机变量，又可以接收相机变量 这里的usb就提现了多态	多态特性是通过接口实现的
	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	//先创建结构体变量
	computer := Computer{}
	phone :=Phone{}
	camera := Camera{}
	//关键点
	computer.Working(phone)
	computer.Working(camera)	//引用传值

	var stu Stu
	var a Ainterface = stu		//如果实现Ainterface,就需要将Binterface，Cinterface的方法都实现，缺一个都会报错
	a.test01()

	var t T = stu	//可以的
	//等同于 var t interface{} = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num float64 = 8.8
	t2 = num	//任何一个变量都可以赋给空接口
	fmt.Println(t2)
}
type Binterface interface {
	test01()
}

type Cinterface interface {
	test02()
}
type Ainterface interface {
	Binterface
	Cinterface
	test03()
}
type Stu struct {}
// 必须要有
//func (stu *Stu) test01() {	//这样写会报错stu类型没有实现Ainterface接口，这里是指针实现的,要调取需要stu类型加上&
func (stu Stu) test01() {
	fmt.Println("test01()")
}
// 必须要有
func (stu Stu) test02() {}
// 必须要有
func (stu Stu) test03() {}

type T interface {}