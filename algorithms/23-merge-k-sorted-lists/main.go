package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 使用归并排序进行两两排序
 */
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	var leftList *ListNode
	var rightList *ListNode
	if length == 2 {
		leftList = lists[0]
		rightList = lists[1]
	} else {
		mid := length / 2
		leftList = mergeKLists(lists[:mid])
		rightList = mergeKLists(lists[mid:])
	}
	node := &ListNode{}
	head := node
	for leftList != nil && rightList != nil {
		node.Next = &ListNode{}
		node = node.Next
		if leftList.Val < rightList.Val {
			node.Val = leftList.Val
			leftList = leftList.Next
		} else {
			node.Val = rightList.Val
			rightList = rightList.Next
		}
	}
	for leftList != nil {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = leftList.Val
		leftList = leftList.Next
	}
	for rightList != nil {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = rightList.Val
		rightList = rightList.Next
	}
	// 因为head头一个初始化时为空
	return head.Next
}
