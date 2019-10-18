package main

func getRow(rowIndex int) []int {
	var results [][]int
	if rowIndex == 0 {
		return []int{1}
	}
	for i := 0; i < rowIndex+1; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, results[i-1][j-1]+results[i-1][j])
			}
		}
		results = append(results, row)
	}
	return results[rowIndex]
}

func optimizedGetRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 0; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
