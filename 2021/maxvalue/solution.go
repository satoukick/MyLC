package main

import (
	"fmt"
	"strconv"
)

func maxValue(n string, x int) string {
	negative := n[0] == '-'
	if negative {
		n = n[1:]
	}

	// result := ""

	if !negative {
		for i := 0; i < len(n); i++ {
			numN, _ := strconv.Atoi(string(n[i]))
			if x > numN {
				return n[:i] + fmt.Sprint(x) + n[i:]
			}
		}
		return n + fmt.Sprint(x)
	}

	for i := 0; i < len(n); i++ {
		numN, _ := strconv.Atoi(string(n[i]))
		if numN > x {
			if i == 0 {
				return string('-') + fmt.Sprint(x) + n
			} else {
				return string('-') + n[:i] + fmt.Sprint(x) + n[i:]
			}
		}
	}
	return string('-') + n + fmt.Sprint(x)

}

func main() {
	maxValue("-13", 2)
}
