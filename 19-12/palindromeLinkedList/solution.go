package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := &ListNode{Next: slow}
	revCur := slow
	for revCur.Next != nil {
		node.Next = revCur.Next
		revCur.Next = revCur.Next.Next
		node.Next.Next = slow
		slow = node.Next
	}

	first := head
	second := slow

	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}
