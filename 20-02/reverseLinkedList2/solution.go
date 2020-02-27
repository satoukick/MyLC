package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	start := pre.Next

	for i := 0; i < n-m; i++ {
		tmp := start.Next
		start.Next = start.Next.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return dummy.Next
}
