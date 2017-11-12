class Solution:
    def pivotIndex(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return -1
        pivot = -1
        leftSum = 0
        rightSum = sum(nums) - nums[0]
        if leftSum == rightSum:
            return 0
        for i in range(1,len(nums)):
            leftSum += nums[i-1]
            rightSum -= nums[i]
            if leftSum == rightSum:
                return i
        return pivot

Solu = Solution()
nums = [-1,-1,-1,0,1,1]
res = Solu.pivotIndex(nums)
print(res)