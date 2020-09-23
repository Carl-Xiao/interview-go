package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		// 1.（保存一下前进方向）保存下一跳
		temp := cur.Next
		// 2.斩断过去,不忘前事
		cur.Next = pre
		// 3.前驱指针的使命在上面已经完成，这里需要更新前驱指针
		pre = cur
		// 当前指针的使命已经完成，需要继续前进了
		cur = temp
	}
	return pre
}

func Print(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Value)
		cur = cur.Next
	}
}

func main() {
	node1 := &ListNode{
		Value: 1,
	}
	node2 := &ListNode{
		Value: 2,
	}
	node3 := &ListNode{
		Value: 3,
	}
	node1.Next = node2
	node2.Next = node3
	Print(node1)
	c := Reverse(node1)
	Print(c)
}
