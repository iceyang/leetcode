package main

import (
	"fmt"
	"testing"
)

func printList(l *ListNode) {
	if l == nil {
		fmt.Println("nil")
		return
	}
	fmt.Print(l.Val, "->")
	printList(l.Next)
}

func buildList(numbers []int) *ListNode {
	node := &ListNode{
		numbers[0],
		nil,
	}
	head := node
	for _, number := range numbers[1:] {
		node.Next = &ListNode{
			Val:  number,
			Next: nil,
		}
		node = node.Next
	}
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	printList(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 2))
	printList(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 1))
	printList(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 5))
}
