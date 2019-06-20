package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := &ListNode{Next: head}
	pre, cur := h, h.Next
	for cur != nil {
		node := cur.Next
		if node == nil {
			return h.Next
		}
		pre.Next = node
		cur.Next = node.Next
		node.Next = cur
		pre = cur
		cur = cur.Next
	}
	return h.Next
}
