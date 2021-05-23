package main

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	counts := make([]int, len(nums))
	for i := 0; i < len(counts); i++ {
		counts[i] = 1
	}

	dp[0] = 1

	maxLength := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					counts[i] = counts[j]
				} else if dp[j]+1 == dp[i] {
					counts[i] += counts[j]
				}
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	result := 0
	for i, n := range dp {
		if n == maxLength {
			result += counts[i]
		}
	}
	return result

	// return countMap[maxLength]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
