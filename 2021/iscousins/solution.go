package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)

func isCousins(root *TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}

	l := []*TreeNode{root}
	for len(l) > 0 {
		newL := []*TreeNode{}

		xFound, yFound := false, false

		for _, node := range l {
			curXFound, curYFound := false, false
			if node.Left != nil {
				newL = append(newL, node.Left)
				if node.Left.Val == x {
					curXFound = true
				}
				if node.Left.Val == y {
					curYFound = true
				}
			}
			if node.Right != nil {
				newL = append(newL, node.Right)
				if node.Right.Val == x {
					curXFound = true
				}
				if node.Right.Val == y {
					curYFound = true
				}
			}
			if curXFound && curYFound {
				return false
			}

			if curXFound {
				xFound = true
			}
			if curYFound {
				yFound = true
			}
		}

		if xFound && yFound {
			return true
		}

		if (xFound && !yFound) || (!xFound && yFound) {
			return false
		}

		l = newL
	}
	return false
}
