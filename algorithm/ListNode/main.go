package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

/**
翻转链表
*/
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

/**
打印数据
*/
func Print(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Value)
		cur = cur.Next
	}
}

//利用数组实现链表的翻转
func Stack(head *ListNode, group int) *ListNode {
	top := head
	var stackNode []*ListNode
	var resultNode []*ListNode
	for top != nil {
		stackNode = append(stackNode, top)
		top = top.Next
		if len(stackNode) == group {
			for i := 0; i < len(stackNode)/2; i++ {
				stackNode[i], stackNode[len(stackNode)-i-1] = stackNode[len(stackNode)-i-1], stackNode[i]
			}
			resultNode = append(resultNode, stackNode...)
			stackNode = nil
		}
	}
	resultNode = append(resultNode, stackNode...)
	resultHead := resultNode[0]
	i := 0
	for i < len(resultNode)-1 {
		resultNode[i].Next = resultNode[i+1]
		i++
	}
	return resultHead
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

	node4 := &ListNode{
		Value: 4,
	}

	node5 := &ListNode{
		Value: 5,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	Print(Stack(node1, 3))
}
