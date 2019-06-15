package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var arr []*ListNode
	length := 0
	node := head
	for node != nil {
		arr = append(arr, node)
		node = node.Next
		length++
	}
	arr = append(arr, nil)
	if length == n {
		return head.Next
	}
	i := length - n
	pre := arr[i-1]
	pre.Next = arr[i+1]
	return head
}
