package main

func getXORSum(arr1 []int, arr2 []int) int {
	result1 := 0
	result2 := 0

	for _, x := range arr1 {
		result1 ^= x
		// 	for _, y := range arr2 {
		// 		result ^= x & y
		// 	}
	}
	for _, y := range arr2 {
		result2 ^= y
	}

	return result1 & result2
}
