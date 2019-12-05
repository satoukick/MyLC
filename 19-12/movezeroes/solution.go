package main

import "fmt"

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := [][]int{
		[]int{0, 1, 0, 3, 12},
		[]int{1, 2, 3, 4, 0},
	}
	for i := 0; i < len(nums); i++ {
		moveZeroes(nums[i])
		fmt.Println(nums[i])
	}
}
