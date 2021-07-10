package main

import "math"

func countTriples(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			cur := math.Sqrt(float64(i*i + j*j))
			if cur == math.Floor(cur) && cur <= float64(n) {
				ans += 2
			}
		}
	}
	return ans
}
