package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col, row := len(matrix[0])-1, 0
	for col >= 0 && row < len(matrix) {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		}
	}
	return false
}
