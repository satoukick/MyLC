package main

import "sort"

func maxSumDivThree(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%3 == 0 {
		return sum
	}
	for length := 1; length < len(nums)-1; length++ {
		for i := 0; i < len(nums)-length; i++ {

		}
	}
	for i := 0; i < len(nums); i++ {
		if sum%3 == 0 {
			return sum
		}
		sum -= nums[i]
	}
	return 0
}
