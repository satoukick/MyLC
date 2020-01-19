package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	numStr := strconv.Itoa(num)

	changed := numStr
	for i := 0; i < len(changed); i++ {
		if changed[i] == '6' {
			changed = changed[:i] + string('9') + changed[i+1:]
			break
		}
	}
	newNum, _ := strconv.Atoi(changed)
	return newNum
}

func main() {
	fmt.Println(maximum69Number(69696))
}
