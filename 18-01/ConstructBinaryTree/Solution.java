package ConstructBinaryTree;

import java.util.HashMap;
import java.util.Map;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        Map<Integer, Integer> inMap = new HashMap<Integer, Integer>();
        for(int i = 0; i < inorder.length; i++)
            inMap.put(inorder[i],i);

        return helper(preorder,inorder,0,preorder.length -1,0,inorder.length - 1,inMap);

    }

    public TreeNode helper(int[] preorder, int[] inorder, int preStart, int preEnd, int inStart, int inEnd,
                           Map<Integer, Integer> inMap) {
        if (preStart > preEnd || inStart > inEnd) return null;

        TreeNode root = new TreeNode(preorder[preStart]);
        int inRoot = inMap.get(root.val);
        int numsLeft = inRoot - inStart;

        root.left = helper(preorder, inorder, preStart + 1, preStart+ numsLeft, inStart, inRoot-1,
                inMap);
        root.right = helper(preorder, inorder, preStart + numsLeft + 1, preEnd, inRoot + 1, inEnd, inMap);

        return root;
    }
}
