package main

func addRungs(rungs []int, dist int) int {
	count := 0

	// pre := make([]int, len(rungs)+1)
	// for i := 1; i <= len(rungs); i++ {
	// 	pre[i] = pre[i-1] + rungs[i-1]
	// 	if pre[i-1]+dist < pre[i] {
	// 		count += rungs[i-1] / dist
	// 	}
	// }
	// sum := make([]int, len(rungs))
	need := 0
	cur := 0

	for i, rung := range rungs {
		need = rung
		var newCount int
		if cur+dist < need {
			if i > 0 {
				// fmt.Println(rungs[i], rungs[i-1], cur, need)
				newCount = (rungs[i] - rungs[i-1]) / dist

				if (rungs[i]-rungs[i-1])%dist == 0 {
					newCount--
				}
			} else {
				newCount = rungs[i] / dist
				if rungs[i]%dist == 0 {
					newCount--
				}
			}
		}
		count += newCount
		cur = need
	}

	return count
}
