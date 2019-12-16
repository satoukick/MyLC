package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	levelNodes := []*TreeNode{root}
	for len(levelNodes) > 0 {
		depth++
		nextLevelNodes := []*TreeNode{}
		for _, n := range levelNodes {
			if n.Left != nil {
				nextLevelNodes = append(nextLevelNodes, n.Left)
			}
			if n.Right != nil {
				nextLevelNodes = append(nextLevelNodes, n.Right)
			}
		}
		levelNodes = nextLevelNodes
	}
	return depth
}
