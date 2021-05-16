package main

import "fmt"

func rotateTheBox(box [][]byte) [][]byte {
	row := len(box)
	col := len(box[0])

	result := make([][]byte, col)
	for i := 0; i < col; i++ {
		result[i] = make([]byte, row)
	}

	for r, line := range box {
		// num of stones before object
		count := 0
		start := 0
		resultCol := row - r - 1
		for i, point := range line {
			if point == '.' {
				continue
			}
			if point == '*' {
				result[i][resultCol] = '*'

				j := i - 1
				for count > 0 {
					// line[j] = '#'
					result[j][resultCol] = '#'
					j--
					count--
				}

				for z := start; z <= j; z++ {
					// line[z] = '.'
					result[z][resultCol] = '.'
				}
				start = i + 1

				continue
			}
			count++
		}

		j := col - 1
		for count > 0 {
			result[j][resultCol] = '#'
			count--
			j--
		}
		for z := start; z <= j; z++ {
			// line[z] = '.'
			result[z][resultCol] = '.'
		}
	}

	for _, r := range result {
		for _, v := range r {
			fmt.Print(string(v))
		}
		fmt.Println()

	}
	fmt.Println("============")
	return result
}

func main() {
	input := [][]byte{}
	r := rotateTheBox(input)
	fmt.Println(r)
}
