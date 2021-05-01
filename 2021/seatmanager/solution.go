package seatmanager

import "container/heap"

type seats []int

func (s seats) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s seats) Len() int {
	return len(s)
}
func (s seats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *seats) Push(x interface{}) {
	num := x.(int)
	*s = append(*s, num)
}

func (s *seats) Pop() interface{} {
	old := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return old
}

type SeatManager struct {
	s *seats
}

func Constructor(n int) SeatManager {
	s := &seats{}
	for i := 0; i < n; i++ {
		(*s) = append((*s), i+1)
	}
	return SeatManager{
		s: s,
	}
}

func (this *SeatManager) Reserve() int {
	item := heap.Pop(this.s)
	return item.(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.s, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
