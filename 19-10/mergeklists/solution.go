package main

import (
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// tad dumb!
func mergeKLists(lists []*ListNode) *ListNode {
	head, cur := &ListNode{}, &ListNode{}
	done := false
	for !done {
		done = true
		min := math.MaxInt32
		minIndex := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				minIndex = i
			}
		}

		if minIndex != -1 {
			done = false
			if head.Next == nil {
				head.Next = lists[minIndex]
			}
			cur.Next = lists[minIndex]
			lists[minIndex] = lists[minIndex].Next
			cur = cur.Next
		}
	}
	return head.Next
}

func main() {

}
