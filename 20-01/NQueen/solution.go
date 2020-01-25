package main

func totalNQueens(n int) int {
	var count int
	columns := make([]bool, n)
	dia1, dia2 := make([]bool, 2*n), make([]bool, 2*n)
	backtracking(0, n, columns, dia1, dia2, &count)
	return count
}

func backtracking(curRow int, n int, columns, d1, d2 []bool, count *int) {
	if curRow == n {
		*count++
		return
	}

	for i := 0; i < n; i++ {
		idx1 := n - i + curRow
		idx2 := curRow + i
		if columns[i] || d1[idx1] || d2[idx2] {
			continue
		}
		columns[i], d1[idx1], d2[idx2] = true, true, true
		backtracking(curRow+1, n, columns, d1, d2, count)
		columns[i], d1[idx1], d2[idx2] = false, false, false
	}
}
