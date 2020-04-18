package main

func minCount(coins []int) int {
	times := 0
	for _, n := range coins {
		times += n / 2
		if n%2 == 1 {
			times++
		}
	}
	return times
}
