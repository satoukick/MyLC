package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	nums := make([]int, num_people)
	c := 0
	index := 0
	for candies > 0 {
		c++
		if c > candies {
			c = candies
		}
		nums[index] += c
		candies -= c
		index = (index + 1) % num_people
	}
	return nums
}

func main() {
	fmt.Println(distributeCandies(10, 3))
}
