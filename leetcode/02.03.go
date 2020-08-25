package main

import (
	"fmt"
	"time"
)

/*
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
*/
func main() {
	st := time.Now()
	//ret := deleteNode("a->b->c->d->e->f", "c")
	t := time.Since(st)
	fmt.Println(ret, t)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(node *ListNode) {
	*node = *node.Next
}

func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
