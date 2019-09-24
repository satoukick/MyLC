package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}
	fizz, buzz := 0, 0
	for i := 1; i <= n; i++ {

		// solution from others 108 ms 93 MB
		fizz++
		buzz++
		val := ""
		if fizz == 3 {
			val += "Fizz"
			fizz = 0
		}
		if buzz == 5 {
			val += "Buzz"
			buzz = 0
		}
		if val == "" {
			val = strconv.Itoa(i)
		}

		// my own solution 108 ms 155 MB
		// if i%3 == 0 {
		// 	val += "Fizz"
		// }
		// if i%5 == 0 {
		// 	val += "Buzz"
		// }
		// if val == "" {
		// 	val = strconv.Itoa(i)
		// }

		res = append(res, val)
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(0))
}
