package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int //数组=>模拟队列
	front   int    //表示指向队列首（不含队首的元素）
	rear    int    //表示指向队列尾
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列已满
	if this.rear == this.maxSize-1 { //！！！重要的提示；rear是队列尾部（含最后元素）
		return errors.New("queue fail")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	if this.front == this.rear {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, nil
}

//显示队列，找到队首然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("加入队列成功！")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Printf("queue.GetQueue() err=%v\n", err)
				return
			}
			fmt.Println("从队列中取出的数据：", val)
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
	/*
		问题：只能加5个数值，无论是否有取出
		对上面代码的小结和说明：./队列-小结和说明.png
	*/
}
