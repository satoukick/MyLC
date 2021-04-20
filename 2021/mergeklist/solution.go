package main

import (
	"container/heap"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkNodeHeap []*ListNode

func (h LinkNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h LinkNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h LinkNodeHeap) Len() int           { return len(h) }
func (h *LinkNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *LinkNodeHeap) Pop() interface{} {
	item := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	if len(lists) == 0 {
		return dummy.Next
	}

	qHeap := &LinkNodeHeap{}

	for _, node := range lists {
		if node != nil {
			heap.Push(qHeap, node)
		}
	}

	pre := dummy

	for qHeap.Len() > 0 {
		nodeItem := heap.Pop(qHeap).(*ListNode)
		pre.Next = nodeItem
		pre = pre.Next
		if nodeItem.Next != nil {
			heap.Push(qHeap, nodeItem.Next)
		}
	}

	return dummy.Next
}
