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

func TestMergeTwoLists(t *testing.T) {
	printList(mergeTwoLists(buildList([]int{1, 2, 4}), buildList([]int{1, 3, 4})))
	printList(mergeTwoLists(nil, buildList([]int{1, 3, 4})))
	printList(mergeTwoLists(buildList([]int{1, 2, 4}), nil))
	printList(mergeTwoLists(nil, nil))
	printList(mergeTwoLists(buildList([]int{1, 2, 4}), buildList([]int{-10, 1, 2, 4})))
}
