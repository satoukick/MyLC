package main

func arraySign(nums []int) int {
	negatives := 0

	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			negatives += 1
		}
	}
	if negatives%2 == 1 {
		return -1
	}
	return 1
}
