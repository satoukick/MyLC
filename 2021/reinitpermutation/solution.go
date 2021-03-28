package main

func reinitializePermutation(n int) int {
	ori := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = i / 2
		} else {
			arr[i] = n/2 + (i-1)/2
		}
		ori[i] = i
	}
	count := 1

	for !compare(ori, arr) {
		count++
		newArr := make([]int, n)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				newArr[i] = arr[i/2]
			} else {
				newArr[i] = arr[n/2+(i-1)/2]
			}
		}
		arr = newArr
	}
	return count
}

func compare(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
