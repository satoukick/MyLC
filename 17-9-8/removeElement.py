# 27. Remove Element
class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        cursor = 0
        for i in range(0,len(nums)):
            if nums[i] != val:
                nums[cursor] = nums[i]
                cursor += 1
        return cursor

