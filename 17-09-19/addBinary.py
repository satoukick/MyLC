class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        if len(a) == 0: return b
        if len(b) == 0: return a

        a = list(reversed(a))
        b = list(reversed(b))
        res = [0] * (max(len(a),len(b)) + 1)
        minlen = min(len(a),len(b))
        for i in range(0,minlen):
            res[i] += int(a[i]) + int(b[i])
            res[i+1] += res[i] / 2
            res[i] = res[i] % 2

        remnant = a if len(a) > len(b) else b

        for i in range(minlen,len(remnant)):
            res[i] += int(remnant[i])
            res[i+1] += res[i] / 2
            res[i] = res[i] % 2

        while len(res) > 1 and res[-1] == 0: res.pop()
        return "".join(map(str,res[::-1]))

res = Solution()
print(res.addBinary('11','1'))