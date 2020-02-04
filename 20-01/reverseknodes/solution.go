package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	cur := head
	var prev *ListNode
	for count < k {
		if cur == nil {
			return head
		}
		count++
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
	reverse(head)
	head.Next = reverseKGroup(cur, k)
	return prev
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := head
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}
	return dummy.Next
}
