package main

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {
	result := make([]int, 0)

	taskSli := []task{}

	for i, t := range tasks {
		tas := task{
			id:             i,
			enqueueTime:    t[0],
			processingTime: t[1],
		}
		taskSli = append(taskSli, tas)
	}

	sort.Slice(taskSli, func(a, b int) bool {
		return taskSli[a].enqueueTime < taskSli[b].enqueueTime
	})

	hp := &hTask{}
	curTime := 0

	var i int

	for i < len(taskSli) {
		// 执行时间最小的任务
		if hp.Len() > 0 {
			t := heap.Pop(hp).(task)
			curTime += t.processingTime
			result = append(result, t.id)
		}

		// 若当前优先队列空，时间进到下一个执行任务的时间
		if hp.Len() == 0 && curTime < taskSli[i].enqueueTime {
			curTime = taskSli[i].enqueueTime
		}

		// 到了执行时间进队列
		for ; i < len(taskSli) && taskSli[i].enqueueTime <= curTime; i++ {
			heap.Push(hp, taskSli[i])
		}
	}

	for hp.Len() > 0 {
		t := heap.Pop(hp).(task)
		result = append(result, t.id)
	}

	return result
}

type task struct {
	enqueueTime    int
	processingTime int
	id             int
}

type hTask []task

func (h hTask) Less(a, b int) bool {
	return h[a].processingTime < h[b].processingTime || (h[a].processingTime == h[b].processingTime && h[a].id < h[b].id)
}

func (h hTask) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h hTask) Len() int {
	return len(h)
}

func (h *hTask) Push(x interface{}) {
	*h = append(*h, x.(task))
}

func (h *hTask) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}
