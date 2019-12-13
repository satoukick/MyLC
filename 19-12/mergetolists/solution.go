package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	var remainNode *ListNode
	if l1 != nil {
		remainNode = l1
	} else {
		remainNode = l2
	}
	for remainNode != nil {
		head.Next = remainNode
		remainNode = remainNode.Next
		head = head.Next
	}
	return dummy.Next
}
