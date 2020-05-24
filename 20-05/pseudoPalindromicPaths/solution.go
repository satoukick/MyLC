func pseudoPalindromicPaths(root *TreeNode) int {
	m := map[int]int{
		root.Val: 1,
	}
	var count int
	search(root, &count, m)
	return count
}

func search(node *TreeNode, count *int, cur map[int]int) {
	if node.Left == nil && node.Right == nil {
		if len(cur) <= 1 {
			*count++
			return
		}
	}

	if node.Left != nil {
		newMap := make(map[int]int)
		for k, v := range cur {
			newMap[k] = v
		}
		if newMap[node.Left.Val] == 1 {
			delete(newMap, node.Left.Val)
		} else {
			newMap[node.Left.Val] = 1
		}
		search(node.Left, count, newMap)
	}

	if node.Right != nil {
		newMap := make(map[int]int)
		for k, v := range cur {
			newMap[k] = v
		}
		if newMap[node.Right.Val] == 1 {
			delete(newMap, node.Right.Val)
		} else {
			newMap[node.Right.Val] = 1
		}
		search(node.Right, count, newMap)
	}
}