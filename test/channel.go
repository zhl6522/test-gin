package main

import "fmt"

var a []int		// slice
var b chan int	// 需要指定通道中元素的类型
func main() {
	fmt.Println(b)		// nil
	b = make(chan int)	// 通道的初始化，通道必须使用make函数初始化才能使用！
	fmt.Println(b)
}
