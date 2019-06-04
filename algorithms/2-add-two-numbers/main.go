package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	value := 0
	if l1 != nil {
		value += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		value += l2.Val
		l2 = l2.Next
	}
	carry := value / 10
	if carry > 0 {
		if l1 != nil {
			l1.Val += carry
		} else if l2 != nil {
			l2.Val += carry
		} else {
			l1 = &ListNode{
				carry,
				nil,
			}
		}
	}
	node := &ListNode{}
	node.Val = value % 10
	node.Next = addTwoNumbers(l1, l2)
	return node
}
