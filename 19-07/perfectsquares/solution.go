package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	border := int(math.Sqrt(float64(n)))
	var min = math.MinInt32
	var nums []int
	helper(n, border, &nums, &min)
	fmt.Println(nums)
	fmt.Println(min)
	return min
}

// n: current num,
func helper(n, sqrt int, nums *[]int, minLen *int) {
	if n < 0 {
		return
	}
	if n == 0 {
		if len(*nums) < *minLen {
			*minLen = len(*nums)
			return
		}
	}
	if len(*nums) >= *minLen {
		return
	}

	for i := sqrt; i >= 1; i-- {
		*nums = append(*nums, i)
		helper(n-i*i, i, nums, minLen)
		*nums = (*nums)[0 : len(*nums)-1]
	}
}

func main() {
	numSquares(13)
}
