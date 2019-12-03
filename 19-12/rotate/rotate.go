package main

import "fmt"

func rotate(nums []int, k int) {
	index := len(nums) - k%len(nums)
	n := append(nums[index:], nums[:index], ...)
	copy(nums, n)
}
