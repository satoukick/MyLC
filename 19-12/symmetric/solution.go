package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root.Left, root.Right}
	for len(stack) > 0 {
		node1, node2 := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		newNodes := []*TreeNode{node1.Left, node2.Right, node1.Right, node2.Left}
		stack = append(stack, newNodes...)
	}
	return true
}
