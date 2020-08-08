package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	cur := 0
	missingCount := 0
	index := 0
	for missingCount < k {
		cur++
		if arr[index] != cur {
			missingCount++
		} else if index != len(arr)-1 {
			index++
		}
	}
	return cur
}

func main() {
	a := findKthPositive([]int{1, 2, 3, 4}, 2)
	fmt.Println(a)
}
