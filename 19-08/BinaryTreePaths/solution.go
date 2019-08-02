package main

import "fmt"

// Given a binary tree, return all root-to-leaf paths.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}
	helper(root, "", &result)
	return result
}

func helper(node *TreeNode, path string, result *[]string) {
	// if node == nil {
	// 	*result = append(*result, path)
	// 	return
	// }

	if len(path) == 0 {
		path += fmt.Sprintf("%d", node.Val)
	} else {
		path += fmt.Sprintf("->%d", node.Val)
	}

	if node.Left != nil {
		helper(node.Left, path, result)
	}
	if node.Right != nil {
		helper(node.Right, path, result)
	}

	// leaf node
	if node.Left == nil && node.Right == nil {
		*result = append(*result, path)
	}
}
