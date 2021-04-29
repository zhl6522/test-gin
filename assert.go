package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

//类型断言
func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point // ok	空接口可以接收任意类型
	fmt.Println(a)
	var b Point
	//b = a 不可以，类型不匹配，就会报panic(恐慌)，要确保原来的空接口指向的就是断言类型。
	b = a.(Point) //类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错。
	fmt.Println(b)

	var b2 float32 = 1.1
	a = b2
	//类型断言(带检测的)
	if y, ok := a.(float64); ok {
		fmt.Println("convert success")
		fmt.Printf("b2的类型是 %T 值=%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行。。。")
	fmt.Println()

	//类型断言的实践
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"sony"}

	//遍历usbArr
	//Phone还有一个特有的方法Call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb接口声明的方法外，还需要调用Phone 特有方法 Call.	=>类型断言
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

	//fmt.Println(usbArr)

	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int = 5
	var name string = "zhl"
	address := "china"
	n4 := 300
	var stu2 Computer
	//var stu2 struct{}
	//fmt.Printf("%T", &stu2)
	TypeJudge(n1, n2, n3, name, address, n4, stu2, &stu2)
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果Usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言...[注意体会！！！]
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

//声明/定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println(p.name, "手机在工作。。。")
}
func (p Phone) Stop() {
	fmt.Println(p.name, "手机已停止")
}
func (p Phone) Call() {
	fmt.Println(p.name, "手机在打电话。。。")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println(c.name, "相机在工作。。。")
}
func (c Camera) Stop() {
	fmt.Println(c.name, "相机已停止")
}

//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, v := range items {
		index++
		switch v.(type) {
		case bool:
			fmt.Printf("第%d个参数是 bool 类型，值是%v\n", index, v)
		case float32:
			fmt.Printf("第%d个参数是 float32 类型，值是%v\n", index, v)
		case float64:
			fmt.Printf("第%d个参数是 float64 类型，值是%v\n", index, v)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%d个参数是 整数 类型，值是%v\n", index, v)
		case string:
			fmt.Printf("第%d个参数是 string 类型，值是%v\n", index, v)
		case struct{}:
			fmt.Printf("第%d个参数是 结构体 类型，值是%v\n", index, v)
		case *struct{}:
			fmt.Printf("第%d个参数是 指针 类型，值是%v\n", index, v)
		default:
			fmt.Printf("第%d个参数是 类型 不确定，值是%v\n", index, v)
		}
	}
}
