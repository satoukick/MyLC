package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		tmp := head.Next
		// head.Next.Next = head
		head.Next = swapPairs(head.Next.Next)
		tmp.Next = head
		return tmp
	}
	return head
}
