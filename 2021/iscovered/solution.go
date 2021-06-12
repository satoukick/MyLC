package main

type ran [][]int

func (s ran) Less(i, j int) int {
	// return s[i][0] < s[i][1] && s[i][]
}

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		covered := false
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				covered = true
				break
			}
		}
		if !covered {
			return false
		}
	}
	return true
}
