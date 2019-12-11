package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	s := strings.TrimSpace(str)
	if s == "" {
		return 0
	}
	var index int
	var negative bool
	if s[0] == '-' {
		negative = true
		index++
	} else if s[0] == '+' {
		index++
	}
	var n int
	for i := index; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num := int(s[i] - '0')

		if !negative {
			if n*10+num > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			if -(n*10 + num) < math.MinInt32 {
				return math.MinInt32
			}
		}
		n = n*10 + num
	}
	if negative {
		n = -n
	}
	return n
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
