package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{}
	h.Next = head
	pre := h
	from := head
	for from != nil && from.Next != nil {
		node := from.Next
		if from.Val == node.Val {
			for node.Next != nil && node.Val == node.Next.Val {
				node = node.Next
			}
			pre.Next = node.Next
			from = pre.Next
		} else {
			pre = from
			from = from.Next
		}
	}
	return h.Next
}
