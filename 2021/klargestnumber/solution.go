package main

import (
	"container/heap"
)

type hp []string

func (h hp) Less(i, j int) bool {
	return compare(h[i], h[j]) == -1
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Len() int {
	return len(h)
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *hp) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}
func kthLargestNumber(nums []string, k int) string {

	h := &hp{}

	for _, s := range nums {

		if h.Len() < k {
			heap.Push(h, s)
			continue
		}

		if compare(s, (*h)[0]) == 1 {
			heap.Pop(h)
			heap.Push(h, s)
		}
	}

	return (*h)[0]
}

func compare(a, b string) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] > b[i] {
			return 1
		}
		return -1
	}
	return 0
}
