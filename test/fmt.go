package main

import "fmt"

func main() {
	var s string
	/*
	Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
	*/
	fmt.Scan(&s)
	fmt.Println("用户输入内容是：", s)
	var (
		name string
		age int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
