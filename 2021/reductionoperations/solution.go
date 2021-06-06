package main

import "container/heap"

type maxHeap []int

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j] || (h[i] == h[j] && i < j)
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func (h *maxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

type minHeap []int

func (h minHeap) Less(i, j int) bool {
	return h[i] > h[j] || (h[i] == h[j] && i < j)
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func (h *minHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func reductionOperations(nums []int) int {
	mxHeap := &maxHeap{}
	for _, n := range nums {
		heap.Push(mxHeap, n)
	}
	mnHeap := &minHeap{}
	count := 0
	curCount := 0
	var largest int
	// nextLargest := (*mxHeap)[0]

	for mxHeap.Len() > 0 {
		count += curCount
		largest = (*mxHeap)[0]
		for mxHeap.Len() > 0 && (*mxHeap)[0] == largest {
			heap.Push(mnHeap, heap.Pop(mxHeap))
			curCount++
		}

	}
	return count
}
