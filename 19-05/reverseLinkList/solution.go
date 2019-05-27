package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		tmp.Next = nil
		cur.Next = cur.Next.Next
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}
	return dummy.Next
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	var next *ListNode = nil

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
