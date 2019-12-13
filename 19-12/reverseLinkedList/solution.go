package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	cur := head
	for cur.Next != nil {
		dummy.Next = cur.Next
		cur.Next = cur.Next.Next
		dummy.Next.Next = head
		head = dummy.Next
	}
	return dummy.Next
}
