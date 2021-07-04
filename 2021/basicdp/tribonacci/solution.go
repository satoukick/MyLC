package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	first, second, third := 1, 1, 0
	for i := 3; i <= n; i++ {
		first, second, third = first+second+third, first, second
	}
	return first
}
