package main

func firstUniqChar(s string) int {
	rs := make(map[rune]int)
	for _, v := range s {
		rs[v]++
	}
	for i, v := range s {
		if rs[v] == 1 {
			return i
		}
	}
	return -1
}
