package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(l *ListNode) *ListNode {
	var l2 *ListNode
	for l != nil {
		l.Next, l, l2 = l2, l.Next, l
	}
	return l2
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	var l3 *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		node := &ListNode{
			Val:  carry,
			Next: l3,
		}
		if l1 != nil && l2 != nil {
			node.Val = node.Val + l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			node.Val = node.Val + l1.Val
			l1 = l1.Next
		} else {
			node.Val = node.Val + l2.Val
			l2 = l2.Next
		}
		if node.Val >= 10 {
			node.Val = node.Val - 10
			carry = 1
		} else {
			carry = 0
		}
		l3 = node
	}
	if carry == 1 {
		node := &ListNode{
			Val:  1,
			Next: l3,
		}
		l3 = node
	}
	return l3
}
