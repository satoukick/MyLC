# 39. Combination Sum
class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        candidates.sort()
        self.dfs(candidates,target,0,[],res)
        return res



    def dfs(self,nums,target,index,path,res):
        if target == 0:
            res.append(path)
            return
        if target < 0:
            return
        for i in range(index,len(nums)):
            self.dfs(nums,target - nums[i],i,path+[nums[i]],res)