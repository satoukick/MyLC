package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root *TreeNode) []int {
	var result []int
	helper(root, &result)
	return result
}

func helper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	helper(node.Left, result)
	helper(node.Right, result)
	*result = append(*result, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{cur.Val}, result...)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	return result
}
