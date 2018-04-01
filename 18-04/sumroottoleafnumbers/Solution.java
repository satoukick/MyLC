package sumroottoleafnumbers;

 class TreeNode {
   int val;
   TreeNode left;
   TreeNode right;
   TreeNode(int x) { val = x; }
}

class Solution {
    public int sumNumbers(TreeNode root) {
        if (root == null)
            return 0;

        return helper(root, 0);
    }

    public int helper(TreeNode root, int curSum){
        if(root == null)
            return 0;

        if(root.left == null && root.right == null) return curSum*10 + root.val;
        return helper(root.left, curSum*10 + root.val) + helper(root.right, curSum*10 + root.val);

    }
}