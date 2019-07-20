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

func TestReverseKGroup(t *testing.T) {
	printList(reverseKGroup(nil, 2))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 0))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 1))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 2))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 3))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 4))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 5))
	printList(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 6))
	// reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 2)
	// reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 3)
	// reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 4)
}
