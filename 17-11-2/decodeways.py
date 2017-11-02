class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) == 0 :return 0
        if not s: return 0

        dp = [0] * (len(s)+1)
        dp[0] = 1
        if s[0] != '0':
            dp[1] = 1
        for i in range(2,len(s)+1):
            if s[i-1] !='0':
                dp[i] += dp[i-1]
            if s[i-2:i] > "09" and s[i-2:i] < "27":
                dp[i] += dp[i-2]
        return dp[len(s)]
