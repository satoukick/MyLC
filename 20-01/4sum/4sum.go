package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			t := target - nums[i] - nums[j]
			// t = -t
			a, b := j+1, len(nums)-1
			for a < b {
				if nums[a]+nums[b] == t {
					result = append(result, []int{nums[i], nums[j], nums[a], nums[b]})
					a++
					b--
					for a < b && nums[a] == nums[a-1] {
						a++
					}
					for a < b && nums[b] == nums[b+1] {
						b--
					}
				} else if nums[a]+nums[b] > t {
					b--
				} else {
					a++
				}

			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, -5, -2, -2, -4, 0, 1, -2}
	fmt.Println(fourSum(nums, -9))
}
