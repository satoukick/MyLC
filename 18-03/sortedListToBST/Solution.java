package sortedListToBST;

//Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }}


 //Definition for a binary tree node.
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public TreeNode sortedListToBST(ListNode head) {
        if (head == null)
            return null;

        ListNode fast = head;
        ListNode slow = head;
        ListNode tmp = null;

        while(fast.next!=null && fast.next.next!=null){
            fast = fast.next.next;
            tmp = slow;
            slow = slow.next;
        }


        if(tmp != null)
            tmp.next = null;
        else
            head = null;
        TreeNode root = new TreeNode(slow.val);
        root.left = sortedListToBST(head);
        root.right = sortedListToBST(slow.next);
        return root;

    }
}
