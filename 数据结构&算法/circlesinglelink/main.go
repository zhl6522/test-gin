package main

import "fmt"

//定义猫的结构体
type CatNode struct {
	No   int
	Name string
	Next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是第一只猫
	if head.Next == nil {
		head.No = newCatNode.No
		head.Name = newCatNode.Name
		head.Next = head //构成一个环状
		fmt.Println(newCatNode, "加入到环形链表中")
		return
	}

	//定义一个临时变量，帮忙找到环形的最后节点
	temp := head
	for true {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	//加入到链表中
	temp.Next = newCatNode
	newCatNode.Next = head
}

//输出环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.Next == nil {
		fmt.Println("空空如也的环形链表")
		return
	}
	for true {
		fmt.Printf("猫的信息：[id=%d name=%s] ->\n", temp.No, temp.Name)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	//空链表
	if temp.Next == nil {
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}
	//如果只有一个节点
	if temp.Next == head {
		if temp.No == id {
			temp.Next = nil
		}
		return head
	}
	//将helper定位到最后
	for true {
		if helper.Next == head {
			break
		}
		helper = helper.Next
	}
	//多个节点
	flag := true
	for true {
		if temp.Next == head { //如果来到这里，说明比较到最后一个了【最后一个还未比较】
			break
		}
		if temp.No == id { //找到这里直接删除
			if temp == head { //说明删除的是头节点
				head = head.Next
			}
			helper.Next = temp.Next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.Next     //移动【比较】
		helper = helper.Next //移动【一旦找到要删除的节点 helper】
	}
	//这里还要比较一次
	if flag { //如果flag为真，则上面没有删除
		if temp.No == id {
			helper.Next = temp.Next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Println("没有这只猫 no=", id)
		}
	}
	return head
}

func main() {
	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		No:   1,
		Name: "Tom",
	}
	cat2 := &CatNode{
		No:   2,
		Name: "Tom2",
	}
	cat3 := &CatNode{
		No:   3,
		Name: "Tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)
	head = DelCatNode(head, 1)
	fmt.Println()
	ListCircleLink(head)
}
