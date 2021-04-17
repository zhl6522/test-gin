package main

import "fmt"

//全局匿名函数
var (
	Func = func(n1, n2 int) int {
		return n1+n2
	}
)

//累加器
func Addupper() func(int) int {
	var n = 10
	return func(i int) int {
		n = n + i
		return n
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
}
