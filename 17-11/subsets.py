class Solution:
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        paths = []
        self.gen_sub_recur(0,[],paths,nums)
        return paths

    def gen_sub_recur(self,index,sub,subs,nums):
        subs.append(sub)
        for i in range(index ,len(nums)):
            self.gen_sub_recur(i+1 , sub+[nums[i]],subs,nums)


Solu = Solution()
n = [1,2,3]
res = Solu.subsets(n)
print(res)