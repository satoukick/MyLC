package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first, second, result := 2, 1, 0
	for i := 3; i <= n; i++ {
		result = first + second
		second = first
		first = result
	}
	return result
}
