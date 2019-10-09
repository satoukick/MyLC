package main

import "fmt"

func rotate(nums []int, k int) {
	startIndex := len(nums) - (k % len(nums))
	new := append(nums[startIndex:], nums[:startIndex]...)
	copy(nums, new)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 1)
	fmt.Println(nums)
}
