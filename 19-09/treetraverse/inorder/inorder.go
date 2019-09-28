package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root *TreeNode) []int {
	result := []int{}
	helper(root, &result)
	return result
}

func helper(node *TreeNode, order *[]int) {
	if node == nil {
		return
	}
	helper(node.Left, order)
	*order = append(*order, node.Val)
	helper(node.Right, order)
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	cur := root
	var stack []*TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		result = append(result, cur.Val)
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return result
}
