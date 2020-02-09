package main

func check(seatsNow *[][]byte, i, j int) bool {
	if (*seatsNow)[i][j] == '#' {
		return false
	}
	// left
	if j > 0 && (*seatsNow)[i][j-1] == 'c' {
		return false
		if i > 0 && (*seatsNow)[i-1][j-1] == 'c' {
			return false
		}
	}
	// right
	if j < len((*seatsNow)[0])-1 && (*seatsNow)[i][j+1] == 'c' {
		return false
		if i > 0 && (*seatsNow)[i-1][j+1] == 'c' {
			return false
		}
	}
	return true
}

func maxStudents(seats [][]byte) int {
	maxNum := 0
	
	// for i := 0; i < len(seats); i++ {
	// 	maxNum = max(maxNum, )
	// }
}

func helper()

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
