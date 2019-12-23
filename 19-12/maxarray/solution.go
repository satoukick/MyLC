package main

func maxSubArray(nums []int) int {
	maxEnd, maxTotal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxEnd = max(maxEnd+nums[i], nums[i])
		maxTotal = max(maxEnd, maxTotal)
	}
	return maxTotal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
