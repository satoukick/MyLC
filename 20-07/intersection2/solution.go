package main

func intersect(nums1 []int, nums2 []int) []int {
	intersectMap := make(map[int]bool)
	countMap := make(map[int]int)
	for _, v := range nums1 {
		intersectMap[v] = true
		countMap[v]++
	}
	result := []int{}
	for _, v := range nums2 {
		if _, ok := intersectMap[v]; ok {
			result = append(result, v)
			if countMap[v] == 1 {
				delete(countMap, v)
				delete(intersectMap, v)
			} else {
				countMap[v]--
			}
		}
	}
	return result
}
