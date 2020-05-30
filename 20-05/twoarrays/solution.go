package main

import "sort"

func canBeEqual(target []int, arr []int) bool {
	sort.Slice(target, func(i, j int) bool { return target[i] < target[j] })

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
