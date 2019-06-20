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
	fmt.Println("nil")
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

func TestSwapPairs(t *testing.T) {
	printList(nil)
	printList(swapPairs(buildList([]int{1})))
	printList(swapPairs(buildList([]int{1, 2, 3, 4})))
	printList(swapPairs(buildList([]int{1, 2, 3, 4, 5})))
}
