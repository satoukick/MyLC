class Solution:
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        if numRows == 0:
            return []
        result = [[1]]
        for i in range(1, numRows):
            row = [1]
            prev = result[i-1]
            for j in range(1, i):
                row.append(prev[j] + prev[j-1])
            row.append(1)
            result.append(row)
        return result

    def generate2(self, numRows):
        triangle = []

        for num in range(numRows):
            row = [None for _ in range(num + 1)]
            row[0], row[-1] = 1, 1

            for i in range(1, num):
                row[i] = triangle[num - 1][i] + triangle[num - 1][i - 1]
            triangle.append(row)

        return triangle


result = Solution()
print(result.generate(5))
print(result.generate2(5))
