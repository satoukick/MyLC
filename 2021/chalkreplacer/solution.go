package main

func chalkReplacer(chalk []int, k int) int {
	roundSum := 0
	for i := 0; i < len(chalk); i++ {
		roundSum += chalk[i]
	}

	remain := k % roundSum
	for i := 0; i < len(chalk); i++ {
		remain -= chalk[i]
		if remain < 0 {
			return i
		}
	}
	return 0
}
