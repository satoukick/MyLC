package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	return isValid(root.Left, minVal, root.Val) && isValid(root.Right, root.Val, maxVal)
}
