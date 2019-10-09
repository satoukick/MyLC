package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)

	for i, v := range inorder {
		inorderMap[v] = i
	}
	return helper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderMap)
}

func helper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	rootPos := inorderMap[rootVal]
	leftGap := rootPos - inStart

	root.Left = helper(preorder, preStart+1, preStart+leftGap, inorder, inStart, rootPos-1, inorderMap)
	root.Right = helper(preorder, preStart+leftGap+1, preEnd, inorder, rootPos+1, inEnd, inorderMap)
	return root
}
