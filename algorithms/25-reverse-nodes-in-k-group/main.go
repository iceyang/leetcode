package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var list *ListNode
	for head != nil {
		node := head
		head.Next, head = list, head.Next
		list = node
	}
	return list
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	i := 0
	list := head
	for head != nil {
		node := head
		head = head.Next
		i += 1
		if i%k == 0 {
			node.Next = nil
			tail := list
			list = reverse(list)
			tail.Next = reverseKGroup(head, k)
			break
		}
	}
	return list
}
