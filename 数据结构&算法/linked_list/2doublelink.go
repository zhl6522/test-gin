package main

import "fmt"

//定义一个HeroNode
type HeroNode2 struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode2	//这个表示指向前一个节点
	next     *HeroNode2 //这个表示指向下一个节点
}

//给双向链表插入一个节点
//编写第一种插入方法，在单链表的最后加入。【简单】
func InsertHeroNodes(head *HeroNode2, newHeroNode *HeroNode2) {
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
	newHeroNode.pre = temp
}

//编写第二种插入方法，根据no的编号从小到大...【实用性高】
func InsertHeroNodes2(head *HeroNode2, newHeroNode *HeroNode2) {
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
		} else if temp.next.no == newHeroNode.no {
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
		newHeroNode.pre = temp
		newHeroNode.next = temp.next

		if temp.next != nil {
			temp.next.pre = newHeroNode	//这两个有先后顺序
		}
		temp.next = newHeroNode

	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

//修改链表的节点信息
func UpdateHeroNodes(head *HeroNode2, newHeroNode *HeroNode2) {
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

//双向链表删除一个节点
func DelHeroNodes(head *HeroNode2, id int) {
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
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("你要删除的ID不存在")
	}
}

//显示链表的所有节点信息
//这里仍然使用单向链表的显示方式
func ListHeroNodes(head *HeroNode2) {
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
	/*for true {
		fmt.Printf("[%d, %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}*/
	fmt.Println()

}
func ListHeroNodes2(head *HeroNode2) {
	//1、创建一个辅助节点[跑龙套，帮忙]
	temp := head
	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	//2、让temp定位到双向链表的最后节点
	for true {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//3、遍历这个链表
	for true {
		fmt.Printf("[%d, %s , %s]==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()

}

//双向链表的应用：水浒英雄的增删改查
func main() {
	//1、先创建一个默认的头节点
	head := &HeroNode2{}
	//2、创建一个新的HeroNode
	hero1 := &HeroNode2{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode2{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode2{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode2{
		no:       4,
		name:     "吴用12",
		nickname: "智多星",
	}
	hero5 := &HeroNode2{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	//3、加入
	InsertHeroNodes2(head, hero1)
	InsertHeroNodes2(head, hero2)
	InsertHeroNodes2(head, hero3)
	InsertHeroNodes2(head, hero4)
	//4、显示
	ListHeroNodes(head)
	fmt.Println("逆序打印")
	ListHeroNodes2(head)
	//修改
	UpdateHeroNodes(head, hero5)
	ListHeroNodes2(head)
	//删除
	DelHeroNodes(head, 4)
	ListHeroNodes2(head)
}
