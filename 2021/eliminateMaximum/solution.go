package main

import (
	"math"
)

func eliminateMaximum(dist []int, speed []int) int {

	minMap := make(map[int]int)

	pre := make([]int, len(dist)+1)

	for i := 0; i < len(dist); i++ {
		mFloat := math.Ceil(float64(dist[i]) / float64(speed[i]))
		// result = min(result, int(mFloat))
		minMap[int(mFloat)]++
	}

	pre[0] = 0
	for i := 1; i <= len(dist); i++ {
		pre[i] = pre[i-1] + minMap[i]
		if pre[i] > i {
			return i
		}
	}
	return len(dist)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
