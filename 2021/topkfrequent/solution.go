package main

import (
	"container/heap"
)

type item struct {
	count int
	word  string
}

type itemQueue []item

func (q itemQueue) Len() int {
	return len(q)
}

func (q itemQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q itemQueue) Less(i, j int) bool {
	return q[i].count < q[j].count || (q[i].count == q[j].count && q[i].word > q[j].word)
}

func (q *itemQueue) Push(x interface{}) {
	value := x.(item)

	*q = append(*q, value)
}

func (q *itemQueue) Pop() interface{} {
	value := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return value
}

func topKFrequent(words []string, k int) []string {
	itemMap := make(map[string]int)

	for _, w := range words {
		itemMap[w]++
	}

	// fmt.Println(itemMap)

	itQueue := &itemQueue{}

	for w, c := range itemMap {
		it := item{
			word:  w,
			count: c,
		}
		heap.Push(itQueue, it)
		if itQueue.Len() > k {
			heap.Pop(itQueue)
		}
	}

	result := make([]string, k)

	for i := k; i >= 1; i-- {
		itValue := heap.Pop(itQueue).(item)
		result[i-1] = itValue.word
	}
	return result
}
