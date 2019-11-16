package main

import (
	"fmt"
	"math"
	"strconv"
)

func encode(num int) string {
	if num == 0 {
		return ""
	}
	n := 1
	length := 1
	for {
		duration := int(math.Pow(2, float64(length)))
		if n+duration-1 >= num {
			break
		}
		length++
		n += duration
	}

	gap := num - n
	result := ""
	for gap != 0 {
		s := gap % 2
		result = strconv.Itoa(s) + result
		gap = gap / 2
	}
	headZeroNum := length - len(result)
	zStr := ""
	for i := 0; i < headZeroNum; i++ {
		zStr += "0"
	}
	result = zStr + result
	return result
}

func main() {
	fmt.Println(encode(8))
	fmt.Println(encode(9))
	fmt.Println(encode(23))
	fmt.Println(encode(107))
}
