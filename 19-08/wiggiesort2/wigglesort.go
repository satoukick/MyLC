package main

import (
	"sort"
)

type numbers []int

func (a numbers) Less(i, j int) bool { return a[i] < a[j] }
func (a numbers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a numbers) Len() int           { return len(a) }

func wiggleSort(nums []int) {
	tmp := make([]int, len(nums))
	sort.Sort(numbers(nums))
	end := len(nums) - 1
	half := len(nums) / 2
	if len(nums)%2 == 0 {
		half--
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			tmp[i] = nums[half]
			half--
		} else {
			tmp[i] = nums[end]
			end--
		}
	}
	copy(nums, tmp)
	//fmt.Println(nums)
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
}
