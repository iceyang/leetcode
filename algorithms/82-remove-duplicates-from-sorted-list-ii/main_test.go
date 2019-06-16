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

func TestDeleteDuplicates(t *testing.T) {
	printList(deleteDuplicates(buildList([]int{1, 2, 3, 3, 4, 4, 5})))
	printList(deleteDuplicates(buildList([]int{1, 1, 1, 2, 3})))
	printList(deleteDuplicates(buildList([]int{1, 1, 1})))
	printList(deleteDuplicates(buildList([]int{1, 1, 1, 2, 2, 3, 4, 4, 5})))
	printList(deleteDuplicates(buildList([]int{1, 1, 1, 2, 2, 3, 4, 4, 5, 5})))
	printList(deleteDuplicates(buildList([]int{1, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5})))
}
