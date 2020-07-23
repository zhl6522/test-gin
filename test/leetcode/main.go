package main

import "fmt"

// 反转链表
// 输入：1->2->3->4->5->6->null
// 输出：6->5->4->3->2->1->null

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next		// 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Printf("%#v\n", head)
	ret := reverseList(head)
	fmt.Printf("%#v\n", ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
