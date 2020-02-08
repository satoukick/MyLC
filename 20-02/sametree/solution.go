package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue, qQueue := []*TreeNode{p}, []*TreeNode{q}
	for len(pQueue) > 0 && len(qQueue) > 0 {
		curP, curQ := pQueue[0], qQueue[0]
		pQueue = pQueue[1:]
		qQueue = qQueue[1:]
		if !check(curP, curQ) {
			return false
		}
		if curP != nil {
			// if !check(curP.Left, curQ.Left) {
			// 	return false
			// }
			pQueue = append(pQueue, curP.Left)
			qQueue = append(qQueue, curQ.Left)
			pQueue = append(pQueue, curP.Right)
			qQueue = append(qQueue, curQ.Right)
			// if curP.Left != nil {

			// }
			// if curQ.Right != nil {

			// }
		}
	}
	return true
}
