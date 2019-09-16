package main

func intersect(nums1 []int, nums2 []int) []int {
	times := make(map[int]int)
	for _, v := range nums1 {
		times[v]++
	}
	var res []int
	for _, v := range nums2 {
		if times[v] > 0 {
			res = append(res, v)
			times[v]--
		}
	}
	return res
}
