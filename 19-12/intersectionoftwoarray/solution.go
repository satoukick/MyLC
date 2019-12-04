package main

func intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	result := []int{}
	for _, n := range nums1 {
		numMap[n]++
	}
	for _, n := range nums2 {
		if _, ok := numMap[n]; ok {
			result = append(result, n)
			numMap[n]--
			if numMap[n] == 0 {
				delete(numMap, n)
			}
		}
	}
	return result
}
