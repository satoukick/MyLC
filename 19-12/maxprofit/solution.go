package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
