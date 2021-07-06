package main

import (
	"errors"
	"fmt"
)

//使用数组模拟一个栈的使用
type Stack struct {
	MaxTop int    //表示我们栈最大可以存放的个数
	Top    int    //表示栈顶，因为栈顶固定，所以我们直接使用Top
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	//先判断是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//遍历栈，注意需要从栈顶开始遍历
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("array[%d]=%d\n", i, this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	//先取值，再this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当栈顶为-1，表示栈为空
	}
	//入栈
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	stack.Push(7)
	//显示
	stack.List()
	//出栈
	val, err := stack.Pop()
	if err != nil {
		fmt.Printf("err=%s\n", err)
		return
	}
	fmt.Println("出栈数据：", val)	//先入后出(FILO)
	//显示
	stack.List()
	//出栈
	val, _ = stack.Pop()
	fmt.Println("出栈数据：", val)	//先入后出(FILO)
	//显示
	stack.List()
}
