package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	cur := 0
	left, right := 0, 0

	def := math.MaxInt32
	result := def

	for right < len(nums) {
		cur += nums[right]

		for cur-nums[left] >= target && left <= right {
			cur -= nums[left]
			left++
		}
		if cur >= target {
			result = min(result, right-left+1)
		}

		right++
	}

	if result == def {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	result := minSubArrayLen(target, nums)
	fmt.Println(result)
}
