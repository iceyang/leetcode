package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	flatten(left)
	flatten(right)

	if left == nil {
		return
	}

	root.Left = nil
	root.Right = left
	if right == nil {
		return
	}
	for node := root.Right; node != nil; node = node.Right {
		if node.Right == nil {
			node.Right = right
			break
		}
	}
}
