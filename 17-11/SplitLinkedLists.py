# 725. Split Linked List in Parts
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def splitListToParts(self, root, k):
        """
        :type root: ListNode
        :type k: int
        :rtype: List[ListNode]
        """
        length = 0
        head = root
        while head:
            length += 1
            head = head.next
        number_in_part = int(length/ k)
        remain_split_points = length % k

        parts = [[]] * k
        head = root
        prev = None
        for i in range(0,len(parts)):
            if prev:
                prev.next = None
            parts[i] = head
            for i in range(number_in_part):
                prev, head = head, head.next
            if remain_split_points:
                remain_split_points -= 1
                prev, head = head, head.next

        return parts


Solu = Solution
