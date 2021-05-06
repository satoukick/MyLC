package main

func decode(encoded []int, first int) []int {
	result := []int{first}
	last := first

	for _, v := range encoded {
		last ^= v
		result = append(result, last)
	}
	return result
}
