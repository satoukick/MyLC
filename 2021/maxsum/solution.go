package main

import "fmt"

func maxSumMinProduct(nums []int) int {
	sum := make([]int, len(nums)+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	left, right := make([]int, len(nums)), make([]int, len(nums))

	stack := []int{}

	// 找到左边最近的比当前元素大的值的下标
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]

		}
		stack = append(stack, i)
	}
	stack = []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = len(nums)
		} else {
			right[i] = stack[len(stack)-1]

		}
		stack = append(stack, i)
	}

	maxValue := 0
	for i := 0; i < len(nums); i++ {
		l := left[i]
		r := right[i]
		s := (sum[r] - sum[l+1]) * nums[i]

		fmt.Println(left[i], right[i], s)

		maxValue = max(maxValue, s)
	}

	for i := 0; i < len(left); i++ {
	}
	fmt.Println(sum)

	return maxValue % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 2}
	maxSumMinProduct(nums)
}
