package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// T: O(n1 + n2)
// S: O(n1 + n2)
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := make([]int, 0), make([]int, 0)

	helper(root1, &leaves1)
	helper(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func helper(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}

	if root.Left != nil {
		helper(root.Left, leaves)
	}
	if root.Right != nil {
		helper(root.Right, leaves)
	}
}
