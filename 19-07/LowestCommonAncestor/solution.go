package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pRoute := findRoute(root, p, []*TreeNode{})
	qRoute := findRoute(root, q, []*TreeNode{})

	var lsa *TreeNode
	length := len(pRoute)
	if len(qRoute) < len(pRoute) {
		length = len(qRoute)
	}
	for i := 0; i < length; i++ {
		if qRoute[i].Val == pRoute[i].Val {
			lsa = qRoute[i]
		}
	}
	return lsa
}

func findRoute(cur, target *TreeNode, route []*TreeNode) []*TreeNode {
	if cur == nil {
		return nil
	}
	route = append(route, cur)
	if cur.Val == target.Val {
		return route
	}
	leftRoute := findRoute(cur.Left, target, route)
	rightRoute := findRoute(cur.Right, target, route)

	if leftRoute != nil {
		return leftRoute
	}
	if rightRoute != nil {
		return rightRoute
	}
	return nil
}

func main() {
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 1}
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}

	result := lowestCommonAncestor(root, p, q)
	fmt.Println(result)
}
