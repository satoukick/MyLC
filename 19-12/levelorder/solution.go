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

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		level := []int{}
		nextNodes := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			level = append(level, nodes[i].Val)
			if nodes[i].Left != nil {
				nextNodes = append(nextNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nextNodes = append(nextNodes, nodes[i].Right)
			}
		}
		result = append(result, level)
		nodes = nextNodes
	}
	return result
}
