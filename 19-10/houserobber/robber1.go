package main

func rob(nums []int) int {
	memo := make([]int, len(nums)+1)

	memo[0], memo[1] = 0, nums[0]
	for i := 1; i < len(nums); i++ {
		memo[i+1] = max(memo[i], memo[i-1]+nums[i])
	}
	return memo[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
