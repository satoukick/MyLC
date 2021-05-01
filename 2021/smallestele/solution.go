package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	smallest := 1
	arr[0] = 1

	for i := 1; i < len(arr); i++ {
		if abs(arr[i], arr[i-1]) <= 1 {
			if arr[i] > arr[i-1] {
				smallest = arr[i]
			}
			continue
		}
		smallest++
		arr[i] = smallest
	}
	return smallest
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
