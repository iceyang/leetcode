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

func TestRotateRight(t *testing.T) {
	list := &ListNode{5, nil}
	list = &ListNode{4, list}
	list = &ListNode{3, list}
	list = &ListNode{2, list}
	list = &ListNode{1, list}
	printList(list)
	// newList := rotateRight(list, 2)
	// newList := rotateRight(list, 0)
	// newList := rotateRight(list, 5)
	newList := rotateRight(list, 28)
	printList(newList)
	printList(rotateRight(&ListNode{}, 0))
}
