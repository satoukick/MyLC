package main

import "fmt"

func binomial(n, k int) int {
	bi := make([][]int, n)
	for i := 0; i < len(bi); i++ {
		bi[i] = make([]int, i+1)
	}
	bi[0][0] = 1
	for i := 1; i < n; i++ {
		bi[i][0], bi[i][i] = 1, 1
		for j := 1; j < len(bi[i])-1; j++ {
			bi[i][j] = bi[i-1][j-1] + bi[i-1][j]
		}
	}

	return bi[n-1][k-1]
}

func main() {
	fmt.Println(binomial(5, 3))
}
