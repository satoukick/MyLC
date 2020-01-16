package main

import "fmt"

// Given a sorted integer array nums, where the range of elements are in the inclusive range [lower, upper], return its missing ranges.

// Example:

// Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
// Output: ["2", "4->49", "51->74", "76->99"]

func findMissingRanges(nums []int, lower int, upper int) []string {
	result := []string{}
	// if len(nums) == 0 {
	// 	return result
	// }
	nums = append([]int{lower - 1}, nums...)
	nums = append(nums, upper+1)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 2 {
			result = append(result, fmt.Sprintf("%d->%d", nums[i-1]+1, nums[i]-1))
		} else if nums[i]-nums[i-1] == 2 {
			result = append(result, fmt.Sprintf("%d", nums[i-1]+1))
		}
	}
	return result
}

func main() {
	nums := []int{0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, 0, 99))
}
