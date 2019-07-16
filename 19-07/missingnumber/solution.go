package main

import "sort"

type number []int

func (a number) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a number) Less(i, j int) bool { return a[i] < a[j] }
func (a number) Len() int           { return len(a) }

func missingNumber(nums []int) int {
	sort.Sort(number(nums))
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)

}

func sumPractice(nums []int) int {
	length := len(nums)
	sum := (1 + length) * (length) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func xorPractice(nums []int) int {
	result := len(nums)
	for i, v := range nums {
		result ^= i
		result ^= v
	}
	return result
}
