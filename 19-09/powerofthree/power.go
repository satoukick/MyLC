package main

import "math"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	var largestPOT int64
	i := 0
	for {
		cur := math.Pow(3, float64(i))
		if cur > math.MaxInt32 {
			largestPOT = int64(cur)
			break
		}
		i++
	}
	return largestPOT%int64(n) == 0
}
