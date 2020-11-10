package main

import (
	"fmt"
	"sort"
)

type point [][]int

func (p point) Len() int {
	return len(p)
}

func (p point) Less(i, j int) bool {
	return p[i][0]*p[i][0]+p[i][1]*p[i][1] < p[j][0]*p[j][0]+p[j][1]*p[j][1]
}

func (p point) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func addpoint(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, K int) [][]int {
	sort.Sort(point(points))

	return points[:K]
}

func main() {
	points := [][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}
	s := kClosest(points, 2)
	fmt.Println(s)
}
