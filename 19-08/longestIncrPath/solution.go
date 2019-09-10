package main

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}

	max := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			len := dfs(matrix, i, j, &cache)
			if len > max {
				max = len
			}
		}
	}
	return max
}

func dfs(matrix [][]int, i, j int, cache *[][]int) int {
	if (*cache)[i][j] != 0 {
		return (*cache)[i][j]
	}
	max := 1
	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
			continue
		}
		if matrix[x][y] <= matrix[i][j] {
			continue
		}
		length := 1 + dfs(matrix, x, y, cache)
		if length > max {
			max = length
		}
	}
	(*cache)[i][j] = max
	return max
}
