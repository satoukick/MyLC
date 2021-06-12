package main

import "sort"

func isBadVersion(version int) bool

func firstBadVersionStdVersion(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}

// Time Complexity: O(n)
// Space Complexity: O(1)

func firstBadVersionBinarySearch(n int) int {
	left, right := 0, n-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return mid
}
