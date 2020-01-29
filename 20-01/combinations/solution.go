package main

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	for i := 1; i <= n-k+1; i++ {
		helper(n, k, []int{i}, i, &result)
	}
	fmt.Println(result)
	return result
}

func helper(n, k int, curCom []int, curN int, result *[][]int) {
	if len(curCom) == k {
		com := make([]int, k)
		copy(com, curCom)
		*result = append(*result, com)
		return
	}

	if curN <= n {
		for i := curN + 1; i <= n; i++ {
			curCom = append(curCom, i)
			fmt.Println(curCom)
			helper(n, k, curCom, i, result)
			curCom = curCom[:len(curCom)-1]
		}
	}
}

func main() {
	combine(4, 2)
}
