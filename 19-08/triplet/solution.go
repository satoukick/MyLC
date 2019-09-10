package main

func increasingTriplet(nums []int) bool {
	x1,x2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= x1 {
			x1 = v
		} else if v <= x2 {
			x2 = v
		} else {
			return true
		}
	}
	return false
}

