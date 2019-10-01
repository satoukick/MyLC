package mains

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	return helper(n1.Left, n2.Right) && helper(n1.Right, n2.Left)
}

func incursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) >= 2 {
		n1 := queue[0]
		n2 := queue[1]
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, []*TreeNode{n1.Left, n2.Right, n1.Right, n2.Left}...)
	}
	return true
}
