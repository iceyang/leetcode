package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	if length == 0 || k%length == 0 {
		return head
	}
	moveStep := length - k%length
	// fmt.Printf("k: %d, length: %d, moveStep: %d\n", k, length, moveStep)
	oldHead := head
	newHead := head
	newTail := head
	for i := 0; i < moveStep; i++ {
		newTail = newHead
		newHead = oldHead.Next
		oldHead = oldHead.Next
	}
	newTail.Next = nil
	tailNode := newHead
	for j := 1; j < k%length; j++ {
		tailNode = tailNode.Next
	}
	tailNode.Next = head
	return newHead
}
