package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	res, first, second := 0, 2, 1
	for i := 3; i <= n; i++ {
		res = first + second
		second = first
		first = res
	}
	return res
}
