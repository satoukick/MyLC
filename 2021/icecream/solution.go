package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	index := 0

	for coins > 0 {
		if index == len(costs) {
			break
		}
		if coins-costs[index] < 0 {
			break
		}
		coins -= costs[index]
		index++
	}
	return index
}
