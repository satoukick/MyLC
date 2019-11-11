package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur, front := head, dummy
	for cur != nil {
		if cur.Val == val {
			front.Next = cur.Next
		} else {
			front = front.Next
		}
		cur = cur.Next
	}

	return dummy.Next
}
