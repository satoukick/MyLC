# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderTraversal(self, root):  # 66ms
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        res = []
        self.helper(root,res)
        return res

    def helper(self,node,res):
        if node:
            self.helper(node.left,res)
            res.append(node.val)
            self.helper(node.right,res)

    def iterative_inorder_traversal(self,root):  # 59ms
        res, stack = [], []
        while True:
            while root:
                stack.append(root)
                root = root.left
            if not stack:
                return res
            node = stack.pop()
            res.append(node.val)
            root = node.right

