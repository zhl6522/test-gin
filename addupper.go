package main

import (
	"fmt"
	"strings"
)

//全局匿名函数
var (
	Func = func(n1, n2 int) int {
		return n1+n2
	}
)

//累加器
func Addupper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(i int) int {
		n = n + i
		str += string(35)	// 35转成成ascII码对应的#
		fmt.Println("str=", str)	//1、str="hell#" 2、str="hello##" 3、str="hello###"
		return n
	}
	//返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n行成一个整体，构成闭包。
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name+suffix
		}
		return name
	}
}

func main() {
	res := func(x, y int) int {
		return x+y
	}(10, 20)
	res2 := func(x, y int) int {
		return x+y
	}
	fmt.Println("res=", res)
	fmt.Println("res2=", res2(20, 30))
	//全局匿名函数的使用
	fmt.Println("res3=", Func(30, 40))
	//使用前面的代码
	f := Addupper()
	fmt.Println(f(1))	//11
	fmt.Println(f(2))	//13
	fmt.Println(f(3))	//16
	//可以这么理解：闭包是类，函数是操作，n是字段。函数和它使用到n构成闭包。
	//当我们反复的调用f函数时，因为n只是初始化一次，因此每调用一次就进行累计。
	//我们要搞清楚闭包的关键，就是要分析出返回的函数它使用（引用）到哪些变量，因为函数和它引用到的变量共同构成闭包。

	//测试makeSuffix的使用
	//返回一个闭包
	//使用闭包的好处：传统方法需要每次都传入相应的条件内容，而闭包因为可以保留上次引用的条件内容，所以我们传入一次就可以反复使用。
	fc := makeSuffix(".jpg")
	fmt.Println("文件名处理后", fc("fire.jpg"))		//fire.jpg
	fmt.Println("文件名处理后", fc("winter"))		//winter.jpg
}
