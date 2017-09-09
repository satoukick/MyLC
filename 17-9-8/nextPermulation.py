class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        i = 0
        j = 0
        for a in range(len(nums) - 1):
            if nums[a] >= nums[a + 1]:
                i = a
                break
            if a == len(nums) - 2:
                return self.reverse(nums, 0, len(nums) - 1)

        for a in range(len(nums) - 1, i - 1, -1):
            if nums[a] > nums[i - 1]:
                j = a
                break

        nums[j], nums[i - 1] = nums[i - 1], nums[j]
        return self.reverse(nums, i, len(nums) - 1)

    def reverse(self, list, l, r):
        while l < r:
            list[l], list[r] = list[r], list[l]
            l += 1
            r -= 1



    def reverse(self,list,l,r):
        while l<r:
            list[l],list[r] = list[r],list[l]
            l += 1
            r -= 1

