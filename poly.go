package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	name	string
}
func (p Phone) Start() {
	fmt.Println("手机在工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机已停止")
}
type Camera struct {
	name	string
}
func (c Camera) Start() {
	fmt.Println("相机在工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机已停止")
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Cameta的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"sony"}
	fmt.Println(usbArr)
}
