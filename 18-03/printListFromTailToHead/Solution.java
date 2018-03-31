package printListFromTailToHead;
import java.util.ArrayList;

class ListNode {
       int val;
       ListNode next = null;

       ListNode(int val) {
           this.val = val;
       }
   }



public class Solution {
    public ArrayList<Integer> printListFromTailToHead(ListNode listNode) {
        ArrayList<Integer> list = new ArrayList<Integer>();
        if (listNode == null)
            return list;
        while (listNode.next != null) {
            list.add(0, listNode.val);
            listNode = listNode.next;
        }
        list.add(0, listNode.val);
        return list;
    }
}