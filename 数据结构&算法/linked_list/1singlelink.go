package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //这个表示指向下一个节点
}

//给链表插入一个节点
	//编写第一种插入方法，在单链表的最后加入。【简单】
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1、先找到该链表的最后这个节点
	//2、创建一个辅助节点[跑龙套，帮忙]
	temp := head
	for true {
		if temp.next == nil { //表示找到最后了
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}
	//编写第二种插入方法，根据no的编号从小到大...【实用性高】
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1、先找到适当的节点
	//2、创建一个辅助节点[跑龙套，帮忙]
	temp := head
	flag := true
	//让插入的节点的no和temp的下一个节点的no比较
	for true {
		if temp.next == nil {
			//说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到temp后面
			break
		}  else if temp.next.no == newHeroNode.no {
			//说明链表中已经有这个no，不能插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("插入的英雄编号已经存在，no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

//修改链表的节点信息
func UpdateHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for true {
		if newHeroNode.no == temp.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		temp.nickname = newHeroNode.nickname
		temp.name = newHeroNode.name
	}
}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for true {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("不要删除的ID不存在")
	}
}

//显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	//1、创建一个辅助节点[跑龙套，帮忙]
	temp := head
	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	//2、遍历这个链表
	for true {
		fmt.Printf("[%d, %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()

}

//单链表的应用：水浒英雄的增删改查
func main() {
	//1、先创建一个默认的头节点
	head := &HeroNode{}
	//2、创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "吴用12",
		nickname: "智多星",
	}
	hero5 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	//3、加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero4)
	//4、显示
	ListHeroNode(head)
	//修改
	UpdateHeroNode(head, hero5)
	ListHeroNode(head)
	//删除
	DelHeroNode(head, 2)
	ListHeroNode(head)
}
