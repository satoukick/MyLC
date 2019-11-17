package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	nums []int
}

func Constructor(root *TreeNode) FindElements {
	nums := []int{}
	var curNode *TreeNode
	root.Val = 0
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		curNode = nodeList[0]
		nodeList = nodeList[1:]

		if curNode.Left != nil {
			curNode.Left.Val = curNode.Val*2 + 1
			nums = append(nums, curNode.Left.Val)
			nodeList = append(nodeList, curNode.Left)
		}
		if curNode.Right != nil {
			curNode.Right.Val = curNode.Val*2 + 2
			nums = append(nums, curNode.Right.Val)
			nodeList = append(nodeList, curNode.Right)
		}
	}
	return FindElements{
		nums: nums,
	}
}

func (this *FindElements) Find(target int) bool {
	for _, v := range this.nums {
		if v == target {
			return true
		}
	}
	return false
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
