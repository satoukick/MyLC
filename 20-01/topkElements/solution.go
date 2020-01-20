package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
	count int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	queue := make(PriorityQueue, len(count))
	var i int
	for k, c := range count {
		item := &Item{
			value: k,
			count: c,
			index: i,
		}
		queue[i] = item
		i++
	}
	heap.Init(&queue)
	result := []int{}
	for i := 0; i < k; i++ {
		item := heap.Pop(&queue).(*Item)
		result = append(result, item.value)
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
