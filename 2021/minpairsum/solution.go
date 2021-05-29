package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)

	result := 0

	for i := 0; i < len(nums)/2; i++ {
		result = max(result, nums[i]+nums[len(nums)-1-i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
