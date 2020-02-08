package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	var curSum, count int
	for i := 0; i < k; i++ {
		curSum += arr[i]
	}
	if curSum/k >= threshold {
		count++
	}

	for i := k; i < len(arr); i++ {
		curSum = curSum - arr[i-k] + arr[i]
		if curSum/k >= threshold {
			count++
		}
	}
	return count
}
