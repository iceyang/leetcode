package main

import (
	"fmt"
	"testing"
)

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println("NULL")
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{3, nil}
	l1 = &ListNode{4, l1}
	l1 = &ListNode{2, l1}
	l1 = &ListNode{7, l1}

	l2 := &ListNode{4, nil}
	l2 = &ListNode{6, l2}
	l2 = &ListNode{5, l2}

	l3 := addTwoNumbers(l1, l2)
	printList(l3)
}
