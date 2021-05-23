package main

import (
	"sort"
)

type envelope [][]int

func (e envelope) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e envelope) Less(i, j int) bool {
	return e[i][0] < e[j][0] || (e[i][0] == e[j][0] && e[i][1] < e[j][1])
}

func (e envelope) Len() int {
	return len(e)
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(envelope(envelopes))

	dp := make([]int, len((envelopes)))
	dp[0] = 1

	result := 1

	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}
	return result

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
