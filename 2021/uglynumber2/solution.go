package main

// Space complexity: O(n)
// Time complexity: O(n)
func nthUglyNumber(n int) int {
	uglies := []int{1}

	last2, last2Index := 2, 0
	last3, last3Index := 3, 0
	last5, last5Index := 5, 0

	for len(uglies) < n {
		smallest := min(last2, min(last3, last5))

		if smallest != uglies[len(uglies)-1] {
			uglies = append(uglies, smallest)
		}

		if smallest == last2 {
			last2Index++
			last2 = 2 * uglies[last2Index]
		} else if smallest == last3 {
			last3Index++
			last3 = 3 * uglies[last3Index]
		} else if smallest == last5 {
			last5Index++
			last5 = 5 * uglies[last5Index]
		}

	}
	// fmt.Println(uglies)
	return uglies[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
