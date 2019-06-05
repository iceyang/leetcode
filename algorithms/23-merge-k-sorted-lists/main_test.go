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

func TestMergeKLists(t *testing.T) {
	var input []*ListNode
	input = append(input, buildList([]int{1, 4, 5}))
	input = append(input, buildList([]int{1, 3, 4}))
	input = append(input, buildList([]int{6, 100}))
	input = append(input, buildList([]int{-2, 6}))
	input = append(input, buildList([]int{2, 7, 8, 10}))
	input = append(input, buildList([]int{7, 11, 102}))
	printList(mergeKLists(input))
}
