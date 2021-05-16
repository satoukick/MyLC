package main

func memLeak(memory1 int, memory2 int) []int {
	t := 1

	for {
		if memory1 >= memory2 {
			if memory1 < t {
				break
			}
			memory1 -= t
		} else {
			if memory2 < t {
				break
			}
			memory2 -= t
		}
		t++
	}

	return []int{t, memory1, memory2}
}

func max(m1, m2 int) int {
	if m1 >= m2 {
		return m1
	}
	return m2
}
