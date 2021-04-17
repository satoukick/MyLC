package main

import (
	"math"
)

// func getMaximumXor(nums []int, maximumBit int) []int {
// 	n := 0
// 	maxBit := int(math.Exp2(float64(maximumBit))) - 1

// 	result := []int{}
// 	lenN := len(nums)

// 	for i := 0; i < lenN; i++ {
// 		for j := 0; j < lenN-i; j++ {
// 			n ^= nums[j]
// 		}
// 		result = append(result, n^maxBit)
// 		n = 0
// 	}
// 	// fmt.Println(result)
// 	return result
// }

func getMaximumXor(nums []int, maximumBit int) []int {
	n := 0
	maxBit := int(math.Exp2(float64(maximumBit))) - 1

	lenN := len(nums)
	result := make([]int, lenN)

	for i := 0; i < lenN; i++ {
		n = n ^ nums[i]
		result[lenN-1-i] = n ^ maxBit
	}
	// fmt.Println(result)
	return result
}

func main() {
	nums := []int{2, 3, 4, 7}

	getMaximumXor(nums, 3)
}
