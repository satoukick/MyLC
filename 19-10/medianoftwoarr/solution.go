package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Index, nums2Index := 0, 0
	sortedArr := []int{}
	for i := 1; i <= len(nums1)+len(nums2); i++ {
		if nums1Index == len(nums1) {
			sortedArr = append(sortedArr, nums2[nums2Index])
			nums2Index++
			continue
		}
		if nums2Index == len(nums2) {
			sortedArr = append(sortedArr, nums1[nums1Index])
			nums1Index++
			continue
		}
		if nums1[nums1Index] <= nums2[nums2Index] {
			sortedArr = append(sortedArr, nums1[nums1Index])
			nums1Index++
		} else {
			sortedArr = append(sortedArr, nums2[nums2Index])
			nums2Index++
		}
	}
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(sortedArr[length/2])
	} else {
		return float64(sortedArr[length/2]+sortedArr[length/2-1]) / 2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
