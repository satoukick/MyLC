package main

import "fmt"

// P(n, k) = P(n-1, k) + k* P(n-1, k-1)
func permuation(n, k int) int {
	pe := make([][]int, n)
	for i := 0; i < len(pe); i++ {
		pe[i] = make([]int, i+2)
	}
	pe[0][0], pe[0][1] = 1, 1
	for i := 1; i < len(pe); i++ {
		pe[i][0], pe[i][i+1] = 1, 1
		for j := 1; j < i+1; j++ {
			pe[i][j] = pe[i-1][j] + j*pe[i-1][j-1]
		}
	}
	return pe[n-1][k]
}

func lessSpace(n, k int) int {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i
	}
	return fact[n] / fact[n-k]
}

func main() {
	fmt.Println(permuation(10, 10))
}
