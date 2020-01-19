package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	return helper(root, target)
}

func helper(node *TreeNode, target int) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = helper(node.Left, target)
	node.Right = helper(node.Right, target)
	if node.Left == nil && node.Right == nil && node.Val == target {
		return nil
	}
	return node
}
