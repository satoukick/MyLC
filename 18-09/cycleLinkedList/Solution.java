package cycleLinkedList;

import java.util.HashSet;
import java.util.Set;

//Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) {
        val = x;
        next = null; }
}

public class Solution {
    public boolean hasCycle(ListNode head) {
        Set<ListNode> used = new HashSet<>();
        while (head != null) {
            if (used.contains(head)){
                return true;
            }
            else {
                used.add(head);
            }
            head = head.next;
        }
        return false;
    }
}