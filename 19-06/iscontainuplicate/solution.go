package main

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		} else {
			numMap[v] = true
		}
	}
	return false
}
