class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        paths = []
        for i in range(0,len(nums)):
            self.dfs(nums, paths, i, [])
        return paths

    def dfs(self, nums, paths, index, cur_p):
        cur_p.sort()
        if cur_p not in paths:
            paths.append(cur_p)
        if index == len(nums):
            return

        for i in range(index, len(nums)):
            self.dfs(nums, paths, i+1, cur_p+[nums[index]])


inputs = [4,4,4,1,4]
solu = Solution()
res = solu.subsetsWithDup(inputs)
print(res)
