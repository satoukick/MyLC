package main

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			return true
		}
		numMap[n] = true
	}
	return false
}
