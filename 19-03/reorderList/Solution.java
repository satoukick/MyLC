package reorderList;


// Definition for singly-linked list.
 public class ListNode {
      int val;
     ListNode next;
    ListNode(int x) { val = x; }
 }

class Solution {
    public void reorderList(ListNode head) {
        if (head == null || head.next == null) {
            return;
        }
        // find the middle node
        ListNode p1 = head, p2 = head;
        while (p2.next != null && p2.next.next != null) {
            p1 = p1.next;
            p2 = p2.next.next;
        }

        ListNode middle = p1;
        ListNode helper = p1.next;
        while(helper.next != null) {
            ListNode cur = helper.next;
            helper.next = cur.next;
            cur.next = middle.next;
            middle.next = cur;
        }

        p1 = head;
        p2 = middle.next;
        while(p1 != middle) {
            middle.next = p2.next;
            p2.next = p1.next;
            p1.next = p2;
            p1 = p2.next;
            p2 = middle.next;
        }
    }
}