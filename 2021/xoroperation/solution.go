package main

// T: O(n)
// S: O(1)

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		cur := start + 2*i
		result ^= cur
	}
	return result
}
