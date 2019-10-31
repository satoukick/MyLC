package main

func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	pow := 1
	var t int

	for i := 1; i <= num; i++ {
		if i == pow {
			pow *= 2
			t = 0
		}
		res[i] = res[t] + 1
		t++
	}
	return res
}
