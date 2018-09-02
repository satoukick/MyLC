package cycleLinkedList2;

import java.util.HashSet;
import java.util.Set;

public class Solution {
    public ListNode detectCycle(ListNode head) {
        Set<ListNode> uniqueNodes = new HashSet<>();
        while(head != null) {
            if (uniqueNodes.contains(head)){
                return head;
            }
            else {
                uniqueNodes.add(head);
            }
            head = head.next;
        }
        return null;
    }
}