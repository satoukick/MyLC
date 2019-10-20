package main

import "math"

func checkStraightLine(coordinates [][]int) bool {
	firstPoint := coordinates[0]
	secondPoint := coordinates[1]
	gradient := 0
	if secondPoint[0]-firstPoint[0] == 0 {
		gradient = math.MaxInt32
	} else {
		gradient = (secondPoint[1] - firstPoint[1]) / (secondPoint[0] - firstPoint[0])
	}
	plus := firstPoint[1] - gradient*firstPoint[0]
	for i := 2; i < len(coordinates); i++ {
		if coordinates[i][0]*gradient+plus != coordinates[i][1] {
			return false
		}
	}
	return true
}
