package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func sum(x int,y int)(ret int) {	//命名返回值
	//return x + y
	ret = x + y
	return
	//return ret
}
func f5() (int, string) {	//多返回值
	return 1,"zhl"
}
func f6(x, y int,m,n string,i,j bool) int {		//参数类型简写：连续多个参数的类型一致时，可以将非最后一个参数的类型省略
	return x + y
}
func f7(x string, y ...int) {	//可变长参数
	fmt.Println(x)
	fmt.Println(y)		// y是int类型的切片 []int
}
func deferFun(i int)  {
	fmt.Println("嘿嘿嘿i", i)
}
func deferNum(i int) {
	defer deferFun(i)
}
// defer多用于函数结束之前释放资源（文件句柄、数据库链接、socket连接）
func deferDemo(i int) {
	fmt.Println("Start")
	defer deferFun(2)
	fmt.Println("End")
}
func main() {
	a := sum(1,2)
	fmt.Println(a)
	m, n:=f5()
	//_, n:=f5()
	fmt.Println(m,n)
	f7("zhl", 1,3,5,7)
	defer deferNum(1)	// defer把它后面的语句延迟到函数即将返回的时候在执行
	defer deferDemo(2)	// 一个函数中可以有多个defer语句，多个defer语句按照先进后出（后进先出）的顺序延迟
	fmt.Println("1")
	//panic("----")

	fmt.Println("---------------语句块作用域---------------")
	if i:=10;i<18 {
		fmt.Println("好好上学")
	}
	//fmt.Println(i)
	for j:=0;j<10;j++ {
		fmt.Printf("%v ",j)
	}
	//fmt.Println(j)

	af1 := f1
	fmt.Printf("\n%T\n", af1)
	af2 := f2
	fmt.Printf("%T\n", af2)
	f7 := f51(f2)
	fmt.Printf("%T\n", f7)
	fmt.Println("---------------匿名函数---------------")
	// 匿名函数多用于函数内部
	fun1 := func(x, y int) {
		fmt.Println(x + y)
	}
	fun1(10,20)

	//如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x+y)
		fmt.Println("Hello World!")
	}(100, 200)

	fmt.Println("---------------闭包---------------")
	// 底层原理：
	// 1、函数可以作为返回值
	// 2、函数内部查找变量的顺序，先在自己内部找，找不到外层找
	// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60
	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90

	ret := f13(f12, 100, 200)
	f11(ret)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

	/*
	你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	分配规则如下：
	a. 名字中每包含1个'e'或'E'分1枚金币
	b. 名字中每包含1个'i'或'I'分2枚金币
	c. 名字中每包含1个'o'或'O'分3枚金币
	d: 名字中每包含1个'u'或'U'分4枚金币
	写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	程序结构如下，请实现 ‘dispatchCoin’ 函数
	*/
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for key1, value1 := range distribution {
		fmt.Printf("%s:%d\n",key1,value1)
	}

	c2(c1, "zhl")
	ff := c3()
	fmt.Printf("%T\n", ff)
	fmt.Println(ff(100,200))
	//闭包示例
	fe:= bi(c1, "uu")
	fmt.Printf("%T\n", fe)
	low(fe)
	// n的阶乘
	fmt.Println(f20(5))
	fmt.Println(taijie(6))
}
// 上台阶的面试题=费布拉切数列(反向)
// 已知楼梯有20阶台阶，上楼可以一步上 1 阶，也可以一步上 2 阶，请编写一个函数计算一种有多少上楼梯的方式。
// 1:1
// 2:1、1，2
// 3:1、1、1,1、2,2、1
// 4:1111，112,121,211,22
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
// 递归
// 递归适合处理那种问题相同/问题规模越来越小的场景
// 递归一定要有一个明确的退出条件
func f20(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n*f20(n-1)
}
func dispatchCoin() (left int) {
	/*for _,name := range users {
		var n21 = strings.Split(name, "")
		for _,v21 := range n21 {
			switch v21 {
			case "e","E":
				distribution[name]+=1
				coins -= 1
			case "i","I":
				distribution[name]+=2
				coins -= 2
			case "o", "O":
				distribution[name]+=3
				coins -= 3
			case "u", "U":
				distribution[name]+=4
				coins -= 4
			}
		}
	}
	return coins*/
	for _,name12 := range users {
		for _,v12 := range name12 {
			switch v12 {
			case 'e', 'E':
				distribution[name12]++
				coins--
			case 'i', 'I':
				distribution[name12]+=2
				coins-=2
			case 'o', 'O':
				distribution[name12]+=3
				coins-=3
			case 'u', 'U':
				distribution[name12]+=4
				coins-=4
			}
		}
	}
	left = coins
	return
}
//闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}
func low(f func()) {
	f()
}
func c1(name string) {
	fmt.Println("hello ", name)
}
// 函数作为参数
func c2(f func(string), name string) {
	f(name)
}
// 函数作为返回值
func c3() (func(int, int) int) {
	return func(i int, i2 int) int {
		return i+i2
	}
}
func adder() func(int) int {
	var x int
	//闭包：函数和其外部变量的引用。
	return func(y int) int {
		x += y
		return x
	}
}
func f1() {
	fmt.Printf("af1")
}
func f2() int {
	return 11
}
func ff(a, b int) int {
	return a + b
}
func f51(x func()int) func(int, int) int {
	return ff
}
func f11(f func()) {
	fmt.Println("this is f11")
	f()
}
func f12(x, y int) {
	fmt.Println("this is f12")
	fmt.Println(x + y)
}
func f13(f func(int, int), m, n int) func() {
	f111 := func() {
		fmt.Println(m + n)
		f(m, n)
	}
	return f111
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}