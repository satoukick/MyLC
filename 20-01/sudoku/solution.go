package main

func solveSudoku(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	bt(&board)
}

func bt(board *[][]byte) bool {
	for i := 0; i < len((*board)); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					byteC := byte(c)
					if isvalid(board, i, j, byteC) {
						(*board)[i][j] = byteC
						if bt(board) {
							return true
						}
						(*board)[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isvalid(board *[][]byte, row, col int, value byte) bool {
	squareRow, squareCol := row/3*3, col/3*3
	for i := 0; i < 9; i++ {
		if (*board)[row][i] == value {
			return false
		}
		if (*board)[i][col] == value {
			return false
		}
		if (*board)[squareRow+i/3][squareCol+i%3] == value {
			return false
		}
	}
	return true
}
