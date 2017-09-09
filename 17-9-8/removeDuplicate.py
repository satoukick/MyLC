class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        newTail = 0
        for i in range(1,len(nums)):
            if nums[i]!=nums[newTail]:
                newTail += 1
                nums[newTail] = nums[i]
        return newTail+1