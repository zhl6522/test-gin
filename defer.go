package main

import "fmt"

func sum(n1, n2 int) int {
	//当执行到defer时，暂不执行，会将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	//在defer将语句放入到栈时，也会将相关的值拷贝同时入栈
	defer fmt.Println("ok1 n1=", n1)	//3、ok1 n1=10
	defer fmt.Println("ok2 n2=", n2)	//2、ok2 n2=20
	//添加一句
	n1++	//n1=11
	n2++	//n2=21
	res := n1 + n2	//res=32
	fmt.Println("ok3 res=", res)	//1、ok3 res=32
	return res
}

//defer最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源：关闭文件资源、释放数据库资源、锁资源。

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)	//res=32
}
