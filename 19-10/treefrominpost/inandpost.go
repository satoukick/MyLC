package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return helper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, inorderMap)
}

func helper(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int, inorderMap map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	rootPos := inorderMap[rootVal]

	root := &TreeNode{Val: rootVal}
	leftGap := rootPos - inStart
	root.Left = helper(inorder, inStart, rootPos-1, postorder, postStart, postStart+leftGap-1, inorderMap)
	root.Right = helper(inorder, rootPos+1, inEnd, postorder, postStart+leftGap, postEnd-1, inorderMap)

	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}
