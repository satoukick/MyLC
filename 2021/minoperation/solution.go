package main

func minOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sum := 0
	for i := 1; i < len(nums); i++ {
		gap := nums[i-1] + 1 - nums[i]
		if gap > 0 {
			sum += gap
			nums[i] = nums[i-1] + 1
		}
	}
	return sum
}
