package main

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		target := -nums[i]
		for j < k {
			if nums[j]+nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}

	}
	return result
}

func main() {

}
