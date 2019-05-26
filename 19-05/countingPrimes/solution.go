package main

func countPrimes(n int) int {
	flag := make([]bool, n)

	count := 0

	for i := 2; i < n; i++ {
		if flag[i] == false {
			count++
			j := i
			for i*j < n {
				flag[i*j] = true
				j++
			}
		}
	}
	return count
}
