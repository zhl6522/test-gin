package main

import (
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int    //4
	array   [4]int //数组
	head    int    //指向队列的队首
	tail    int    //指向队列的队尾
}

//入队列	AddQueue(push)	GetQueue(pop)
func (this *CircleQueue) Push(val int) (err error) {
	return nil
}

func (this *CircleQueue) Pop() (val int, err error) {
	return val,nil
}
//判断环形队列是否为空
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}
//判断环形队列是否满了


func main() {
	circle := &CircleQueue{
		maxSize: 4,
		head:    -1,
		tail:    -1,
	}

	var key string
	var val int
	for true {
		fmt.Println("1、输入add 表示添加数据到队列")
		fmt.Println("2、输入get 表示从队列获取数据")
		fmt.Println("3、输入show 表示显示队列")
		fmt.Println("4、输入exit 表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			//fmt.Println("输入你要入队列数")
			//fmt.Scanln(&val)
			//err := queue.AddQueue(val)
			//if err != nil {
			//	fmt.Println(err.Error())
			//	return
			//} else {
			//	fmt.Println("加入队列成功！")
			//}
		case "get":
			//val, err := queue.GetQueue()
			//if err != nil {
			//	fmt.Printf("queue.GetQueue() err=%v\n", err)
			//	return
			//}
			//fmt.Println("从队列中取出的数据：", val)
		case "show":
			//queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
