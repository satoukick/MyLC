package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum := make(map[int]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum[A[i]+B[j]]++
		}
	}

	// fmt.Println(sum)

	count := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			fmt.Println(-(C[i] + D[j]))
			if _, ok := sum[-(C[i] + D[j])]; ok {
				count += sum[-(C[i] + D[j])]
			}
		}
	}
	return count
}

func main() {
	A := []int{-1, -1}
	B := []int{-1, 1}
	C := []int{-1, 1}
	D := []int{1, -1}
	fourSumCount(A, B, C, D)
}

// [-1,-1]
// [-1,1]
// [-1,1]
// [1,-1]
