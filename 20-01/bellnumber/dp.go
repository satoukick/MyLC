package main

import "fmt"

func bellNumber(n int) int {
	bell := make([][]int, n)
	for i := 0; i < len(bell); i++ {
		bell[i] = make([]int, i+1)
	}
	bell[0][0] = 1

	for i := 1; i < len(bell); i++ {
		bell[i][0] = bell[i-1][i-1]
		for j := 1; j < len(bell[i]); j++ {
			bell[i][j] = bell[i][j-1] + bell[i-1][j-1]
		}
	}

	return bell[n-1][n-1]
}

func main() {
	fmt.Println(bellNumber(5))
}
