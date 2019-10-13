package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	dir := []int{-1, 0, 1}
	exist := make([][]bool, 8)
	result := [][]int{}
	for i := 0; i < 8; i++ {
		exist[i] = make([]bool, 8)
	}
	for _, q := range queens {
		exist[q[0]][q[1]] = true
	}
	for _, x := range dir {
		for _, y := range dir {
			if x == 0 && y == 0 {
				continue
			}
			r, c := king[0], king[1]
			for r+x >= 0 && r+x < 8 && c+y >= 0 && c+y < 8 {
				r += x
				c += y
				if exist[r][c] {
					result = append(result, []int{r, c})
					break
				}
			}
		}
	}
	return result
}
