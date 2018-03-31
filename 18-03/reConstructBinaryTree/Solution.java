package reConstructBinaryTree;


import java.util.HashMap;
import java.util.Map;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class Solution {
    public TreeNode reConstructBinaryTree(int [] pre,int [] in) {
        Map<Integer, Integer> inMap = new HashMap<Integer, Integer>();

        for (int i = 0; i < in.length; i++)
             inMap.put(in[i],i);

        return helper(pre,0, pre.length - 1, in, 0, in.length - 1, inMap);
    }

    public TreeNode helper(int[] preorder, int preStart, int preEnd, int[] inorder, int inStart, int inEnd, Map<Integer, Integer> inMap){
        if(preStart > preEnd || inStart > inEnd)
            return null;

        TreeNode root = new TreeNode(preorder[preStart]);
        int inPosition = inMap.get(root.val);
        int numsLeft = inPosition - inStart;

        root.left = helper(preorder,preStart+1, preStart+inPosition,inorder,inStart, inPosition-1, inMap);
        root.right = helper(preorder, preStart + numsLeft + 1, preEnd, inorder, inPosition+1, inEnd, inMap);

        return root;
    }


}