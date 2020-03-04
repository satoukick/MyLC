package main

func orangesRotting(grid [][]int) int {
	type point struct {
		x, y int
	}
	queue := []point{}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			}
		}
	}
	round := 0
	for count > 0 && len(queue) > 0 {
		round++
		n := len(queue)
		for i := 0; i < n; i++ {
			r, c := queue[0].x, queue[0].y
			queue = queue[1:]
			if r > 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				count--
				queue = append(queue, point{r - 1, c})
			}
			if r < len(grid)-1 && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				count--
				queue = append(queue, point{r + 1, c})
			}
			if c > 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				count--
				queue = append(queue, point{r, c - 1})
			}
			if c < len(grid[0])-1 && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				count--
				queue = append(queue, point{r, c + 1})
			}
		}
	}
	if count > 0 {
		return -1
	}
	return round
}
