package main

import "fmt"

func balancedStringSplit(s string) int {
	total := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			j++
		} else if s[i] == 'R' {
			j--
		}

		if j == 0 {
			total++
		}
	}
	return total
}

func main() {
	st := "LRLLLRRLLRRRLLLRRR"
	fmt.Println(balancedStringSplit(st))
}
