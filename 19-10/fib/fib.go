package main

import (
	"fmt"
)

func fib(N int) int {
	if N < 2 {
		return N
	}
	first, second, result := 1, 0, 0
	for i := 2; i <= N; i++ {
		result = first + second
		second = first
		first = result
	}
	return result
}

func main() {
	fmt.Println(fib(3))
}
