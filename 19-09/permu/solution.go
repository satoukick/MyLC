package main

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	helper(&result, []int{}, nums)
	return result
}

func helper(result *[][]int, curPer []int, remain []int) {
	if len(remain) == 0 {
		cur := make([]int, len(curPer))
		copy(cur, curPer)
		*result = append(*result, cur)
		return
	}

	for i := 0; i < len(remain); i++ {
		curPer = append(curPer, remain[i])
		newr := make([]int, len(remain))
		copy(newr, remain)
		newr = append(newr[:i], newr[i+1:]...)
		helper(result, curPer, newr)
		curPer = curPer[:len(curPer)-1]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}
