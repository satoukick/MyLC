package main

func isHappy(n int) bool {
	nums := make(map[int]bool)
	cur := n
	nums[cur] = true

	for {
		var res int
		for cur != 0 {
			n := cur % 10
			res += n * n
			cur = cur / 10
		}

		if res == 1 {
			return true
		}

		if _, ok := nums[res]; ok {
			break
		}
		nums[res] = true
		cur = res
	}
	return false
}
