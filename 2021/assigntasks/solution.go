package main

import (
	"container/heap"
	"fmt"
)

type srv struct {
	weight int
	index  int
}

type serverPQ []*srv

func (p serverPQ) Less(i, j int) bool {
	if p[i].weight < p[j].weight {
		return true
	} else if p[i].weight > p[j].weight {
		return false
	}

	return p[i].index < p[j].index
}

func (p serverPQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p serverPQ) Len() int {
	return len(p)
}

func (p *serverPQ) Push(x interface{}) {
	item := x.(*srv)
	*p = append(*p, item)
}

func (p *serverPQ) Pop() interface{} {
	item := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return item
}

type working struct {
	t     int
	index int
}

type workingPQ []*working

func (p *workingPQ) Push(x interface{}) {
	item := x.(*working)
	*p = append(*p, item)
}

func (p *workingPQ) Pop() interface{} {
	item := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return item
}

func (p workingPQ) Less(i, j int) bool {
	if p[i].t < p[j].t {
		return true
	}

	if p[i].t == p[j].t && p[i].index < p[j].index {
		return true
	}
	return false

}

func (p workingPQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p workingPQ) Len() int {
	return len(p)
}

func assignTasks(servers []int, tasks []int) []int {
	works := new(workingPQ)
	idle := new(serverPQ)

	for i, server := range servers {
		s := &srv{
			weight: server,
			index:  i,
		}
		heap.Push(idle, s)
	}

	result := make([]int, 0)
	curTime := 0

	for i, ti := range tasks {
		if i > curTime {
			curTime = i
		}

		if idle.Len() == 0 {
			curTime = (*works)[0].t
		}

		for works.Len() > 0 {
			workTop := (*works)[0]
			if workTop.t > curTime {
				break
			}

			heap.Pop(works)
			idleServer := &srv{
				index:  workTop.index,
				weight: servers[workTop.index],
			}
			heap.Push(idle, idleServer)
		}

		curServer := heap.Pop(idle).(*srv)
		work := &working{
			t:     ti + curTime,
			index: curServer.index,
		}
		heap.Push(works, work)

		result = append(result, work.index)
	}
	return result
}

func main() {
	servers := []int{3, 3, 2}
	task := []int{1, 2, 3, 2, 1, 2}

	fmt.Println(assignTasks(servers, task))
}
