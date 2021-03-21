package main

import (
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	ans := 0
	for i := 0; i < len(coins); i++ {
		// reach the n-1 of coin
		if ans >= coins[i]-1 {
			ans += coins[i]
			continue
		}
		break
	}
	return ans + 1
}
