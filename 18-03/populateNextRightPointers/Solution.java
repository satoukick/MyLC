package populateNextRightPointers;


class TreeLinkNode {
    int val;
    TreeLinkNode left, right, next;
    TreeLinkNode(int x) { val = x; }
}

public class Solution {
    public void connect(TreeLinkNode root) {
        while(root != null){
            TreeLinkNode pre = root;
            while(pre != null && pre.left != null){
                pre.left.next = pre.right;
                pre.right.next = pre.next == null? null:pre.next.left;
                pre = pre.next;
            }
            root = root.left;
        }
    }
}