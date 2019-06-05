package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for l1 != nil && l2 != nil {
		node.Next = &ListNode{}
		node = node.Next
		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
	}
	for l1 != nil {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = l1.Val
		l1 = l1.Next
	}
	for l2 != nil {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = l2.Val
		l2 = l2.Next
	}
	// 因为head头一个初始化时为空
	return head.Next
}
