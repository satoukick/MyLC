package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	n := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
