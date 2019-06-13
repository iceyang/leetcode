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

func TestReverseBetween(t *testing.T) {
	list := &ListNode{5, nil}
	list = &ListNode{4, list}
	list = &ListNode{3, list}
	list = &ListNode{2, list}
	list = &ListNode{1, list}
	printList(list)
	printList(reverseBetween(list, 1, 2))
}
