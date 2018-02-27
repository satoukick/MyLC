package BinaryTreeZigzag;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

  // Definition for a binary tree node.
class TreeNode {
      int val;
      TreeNode left;
      TreeNode right;
      TreeNode(int x) { val = x; }
 }

class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        travel(root, 0, result);
        return result;

    }

    public void travel(TreeNode currentNode, int level, List<List<Integer>> result){
        if (currentNode == null) return;

        if(result.size() <= level) result.add(new LinkedList<>());

        List<Integer> collection = result.get(level);
        if (level % 2 == 0) collection.add(currentNode.val);
        else collection.add(0,currentNode.val);

        travel(currentNode.left, level + 1, result);
        travel(currentNode.right, level + 1, result);
    }
}