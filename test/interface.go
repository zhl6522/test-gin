package main

import "fmt"

// 引出接口的实例
// 接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法。

// 定义一个能叫的类型
type speaker interface {
	speak()		// 只要实现了speak方法的变量都是speaker类型，方法签名
}
type dog struct {}
type cat struct {}
type person struct {}
func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("汪汪汪~")
}
func (p person) speak() {
	fmt.Println("啊啊啊~")
}
func da(x speaker) {
	// 接受一个参数，传进来什么就打什么
	x.speak()	// 挨打了就会叫
}

// 定义一个car接口类型
// 不管是什么结构体，只要有run方法都是car类型
type car interface {
	run()
}
type falali struct {
	brand	string
}
func (f falali) run() {
	fmt.Printf("%s750迈~\n", f.brand)
}
type baoshijie struct {
	brand	string
}
func (b baoshijie) run() {
	fmt.Printf("%s700迈~\n", b.brand)
}
// drive函数接受一个car类型的变量
func drive(c car) {
	c.run()
}

// 接口的实现
type animal interface {
	move()
	eat(string)
}
type cat2 struct {
	name	string
	feet	int8
}

func (c cat2)move() {
	fmt.Println("走猫步")
}
/*func (c cat2) eat() {		// 与animal的eat(string)实现的不是同一个方法，不满足接口的需求
	fmt.Println("吃猫粮~")
}*/
func (c cat2) eat(food string) {
	fmt.Printf("吃猫%s~\n", food)
}
type chicken struct {
	feet	int8
}
func (c chicken) move() {
	fmt.Println("鸡动")
}
func (c chicken) eat(w string) {
	fmt.Println("%s吃鸡饲料~", w)
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

	var f1 = falali{
		"法拉利",
	}
	var b1 = baoshijie{
		"保时捷",
	}
	drive(f1)
	drive(b1)

	var a1 animal		// 定义一个接口类型的变量
	fmt.Printf("%T\n", a1)		// <nil> 因为它是一个动态类型，动态值
	bc := cat2{
		"淘气",
		4,
	}
	a1 = bc
	a1.eat("猫粮")
	fmt.Printf("%T\n", a1)	// 返回main.cat2 看interface.png详解
	fmt.Println(a1)

	kfc := chicken{
		feet:2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1)	// 返回main.chicken 看interface.png详解

	// 一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以以称为这个接口类型的变量。
	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
