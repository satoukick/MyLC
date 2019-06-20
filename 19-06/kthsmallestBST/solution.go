package main

import "sort"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Number []int

func (a Number) Less(i, j int) bool { return a[i] < a[j] }
func (a Number) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Number) Len() int           { return len(a) }

func kthSmallest(root *TreeNode, k int) int {
	nums := []int{}
	helper(root, &nums)
	sort.Sort(Number(nums))
	return nums[k-1]
}

func helper(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	*nums = append(*nums, node.Val)
	helper(node.Left, nums)
	helper(node.Right, nums)
}
