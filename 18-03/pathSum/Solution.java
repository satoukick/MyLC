package pathSum;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> result = new LinkedList<List<Integer>>();
        List<Integer> currentPath = new LinkedList<Integer>();
        pathSum(root,sum,result,currentPath);
        return result;
    }

    public void pathSum(TreeNode root, int sum, List<List<Integer>> result, List<Integer> currentPath){
        if(root == null)
            return;

        currentPath.add(root.val);
        if(root.left == null && root.right == null && sum == root.val){
            result.add(new LinkedList<Integer>(currentPath));
        }else{
            pathSum(root.left, sum - root.val, result, currentPath);
            pathSum(root.right, sum- root.val, result, currentPath);
        }

        currentPath.remove(currentPath.size() -1);
    }
}