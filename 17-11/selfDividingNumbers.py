class Solution:
    def selfDividingNumbers(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        list = []
        zero = '0'
        for num in range(left,right+1):
            isValid = True
            nstr = str(num)
            if zero in nstr:
                continue
            for n in nstr:
                if num % int(n) != 0:
                    isValid = False
                    break
            if isValid:
                list.append(num)

        return list

Solu = Solution()
res = Solu.selfDividingNumbers(1,22)
print(res)