package main


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	dummy := ListNode{0, head.Next}

	for even.Next != nil && even.Next.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	if even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = nil
	}

	odd.Next = dummy.Next
	return head
}