package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return helper(nums)
}

func helper(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := len(nums) / 2
	root := &TreeNode{Val: nums[index]}
	root.Left = helper(nums[:index])
	root.Right = helper(nums[index+1:])

	return root
}
