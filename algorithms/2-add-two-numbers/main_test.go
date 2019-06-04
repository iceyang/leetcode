package main

import (
	"fmt"
	"testing"
)

func printListNode(l *ListNode) {
	if l == nil {
		return
	}
	printListNode(l.Next)
	fmt.Print(l.Val)
}

func buildList(number int) *ListNode {
	node := &ListNode{}
	if number == 0 {
		node.Val = 0
		node.Next = nil
		return node
	}
	node.Val = number % 10
	if number/10 == 0 {
		node.Next = nil
	} else {
		node.Next = buildList(number / 10)
	}
	return node
}

func TestAddTwoNumbers(t *testing.T) {
	printListNode(addTwoNumbers(buildList(342), buildList(465)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(0), buildList(465)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(342), buildList(0)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(0), buildList(0)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(0), buildList(999)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(999), buildList(999)))
	fmt.Println()
	printListNode(addTwoNumbers(buildList(999), buildList(1000)))
	fmt.Println()
}
