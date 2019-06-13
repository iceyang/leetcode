package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	pre := head
	start := head
	for i := 1; i < m; i++ {
		pre = start
		start = start.Next
	}
	tail := start
	var list *ListNode = nil
	for i := m; i <= n; i++ {
		node := start.Next
		start.Next = list
		list = start
		start = node
	}
	tail.Next = start
	if m == 1 {
		return list
	} else {
		pre.Next = list
	}
	return head
}
