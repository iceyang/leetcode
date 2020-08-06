package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob(a) = Max(
//     rob(a.Left) + rob(a.Right)
//     a.Val + rob(a.Left.Left) + rob(a.Left.Right) + rob(a.Right.Left) + rob(a.Right.Right),
// )
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	left, right := 0, 0
	if root.Left != nil {
		left = rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		right = rob(root.Right.Left) + rob(root.Right.Right)
	}
	return max(
		rob(root.Left)+rob(root.Right),
		root.Val+left+right,
	)
}
