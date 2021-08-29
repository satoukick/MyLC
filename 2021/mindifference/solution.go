package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := math.MaxInt32

	for i := 0; i < len(nums)-k+1; i++ {
		result = min(result, nums[i+k-1]-nums[i])
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
