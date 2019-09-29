package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	li := []*TreeNode{root}
	for len(li) > 0 {
		cur := make([]*TreeNode, len(li))
		copy(cur, li)
		li = []*TreeNode{}
		level := []int{}
		for len(cur) > 0 {
			node := cur[0]
			cur = cur[1:]
			if node.Left != nil {
				li = append(li, node.Left)
			}
			if node.Right != nil {
				li = append(li, node.Right)
			}
			level = append(level, node.Val)
		}
		result = append(result, level)
	}
	return result
}
