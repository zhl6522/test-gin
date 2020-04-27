package series

import "fmt"

func init() {
	fmt.Println("init1")
}
//同医源码文件，可以定义多个init函数
func init() {
	fmt.Println("init2")
}

func Square(n int) int {	//square小写字母的方法不能访问到
	return n*n
}

func GetFibonacci(n int) []int {
	fibList := []int{1,1}
	for i:=2;i<n;i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
