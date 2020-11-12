package main

func sortArrayByParityII(A []int) []int {
	i, j := 0, 1
	for {
		if A[i]%2 == 1 && A[j]%2 == 0 {
			A[j], A[i] = A[i], A[j]
			i += 2
			j += 2
		} else if A[i]%2 == 0 {
			i += 2
		} else if A[j]%2 == 1 {
			j += 2
		}
		if i > len(A)-2 || j > len(A)-1 {
			break
		}
	}
	return A
}
