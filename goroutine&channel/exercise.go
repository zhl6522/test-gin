package main

import "fmt"

type Cat struct {
	Name	string
	Age		int
}

func main() {
	//定义一个存放任意数据类型的管道 3个数据
	//var allChan chan interfave{}
	allChan := make(chan interface{}, 3)
	allChan<-10
	allChan<-"jack"
	allChan<-Cat{Name:"tom",Age:18}

	//我们希望获得管道中第三个元素，则先将前两个推出
	<-allChan
	<-allChan

	first := <-allChan
	fmt.Printf("first")
	//fmt.Println(first.Name)	//这么写类型不匹配，就会报错(first.Name undefined (type interface {} is interface with no methods))，要确保原来的空接口指向的就是断言类型。
	fmt.Println(first.(Cat).Name)
	//fmt.Println(allChan)
}
