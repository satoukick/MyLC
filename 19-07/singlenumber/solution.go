package main

import (
	"fmt"
	"sort"
)

// The order of the result is not important. So in the above example, [5, 3] is also correct.
// Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

type Numbers []int

func (a Numbers) Less(i, j int) bool { return a[i] < a[j] }
func (a Numbers) Len() int           { return len(a) }
func (a Numbers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func singleNumber(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sort.Sort(Numbers(nums))
	index := 1
	for index < len(nums) {
		if nums[index-1] == nums[index] {
			nums = append(nums[0:index], nums[index+1:]...)
			index--
		}
		index++
	}
	return nums
}

func main() {
	fmt.Println(singleNumber([]int{4, 4, 2, 1, 6, 6}))
}
