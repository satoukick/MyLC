class Solution(object):
        def searchMatrix(self, matrix, target):
            """
            :type matrix: List[List[int]]
            :type target: int
            :rtype: bool
            """
            if not matrix or target is None:
                return False

            row, col = len(matrix), len(matrix[0])
            low, high = 0, row * col - 1

            while low <= high:
                mid = (low + high) / 2

                num = matrix[mid / col][mid % col]

                if num == target:
                    return True
                elif num < target:
                    low = mid + 1
                else:
                    high = mid - 1

            return False
