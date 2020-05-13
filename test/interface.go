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
// 使用值接收者实现了接口的所有方法
/*func (c cat2)move() {
	fmt.Println("走猫步")
}
//func (c cat2) eat() {		// 与animal的eat(string)实现的不是同一个方法，不满足接口的需求
//	fmt.Println("吃猫粮~")
//}
func (c cat2) eat(food string) {
	fmt.Printf("吃猫%s~\n", food)
}*/
// 使用指针接收者实现了接口的所有方法
func (c *cat2)move() {
	fmt.Println("走猫步")
}
func (c *cat2) eat(food string) {
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

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal3 interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat()
}
type cat3 struct {
	name	string
	feet	int8
}
// cat3实现了mover接口
func (c *cat3)move() {
	fmt.Println("走猫步")
}
// cat3实现了eater接口
func (c *cat3) eat(food string) {
	fmt.Printf("吃猫%s~\n", food)
}

// 空接口：没有必要起名字，通常定义成下面的格式：
//interface{}
// 所有的类型都是先了空接口，也就是任意类型的变量都能保存空接口中。

// 空接口作为你函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
// 类型断言	我想知道空接口接收值具体是什么
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Print("猜错了\n")
	} else {
		fmt.Printf("传进来的是一个字符串：%v\n", str)
	}
}
// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Printf("传进来的是一个字符串：%v\n", t)
	case int:
		fmt.Printf("传进来的是一个int：%v\n", t)
	case int64:
		fmt.Printf("传进来的是一个int64：%v\n", t)
	case bool:
		fmt.Printf("传进来的是一个bool：%v\n", t)
	}
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
	a1 = &bc
	//a1 = bc
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

	fmt.Println("---------------值接收者与指针接收者---------------")
	// 使值接收者与指针接收者的区别？
	// 1、使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存。
	// 2、指针接收者实现接口只能存结构体指针类型的变量。

	c4 := cat2{"tom", 4}
	c5 := &cat2{"假老练", 4}
	a1 = &c4			// 实现anim这个接口的是cat指针，所以要使用 &c4
	fmt.Println(a1)
	a1 = c5
	fmt.Println(a1)


	fmt.Println("---------------空接口---------------")
	// interface：关键字
	// interface{}：空接口类型
	var i  map[string]interface{}
	i = make(map[string]interface{}, 16)
	i["name"] = "zhl"
	i["age"] = 25
	i["merried"] = false
	i["hobby"] = [...]string{"轮滑","滑雪", "rap"}
	fmt.Println(i)

	show(false)
	show(nil)
	show(i)

	assign(100)
	assign2(true)
	assign2("哈哈哈")
	assign2(int64(200))

}
