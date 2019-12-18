# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        left, right = 1, n
        min_version = 1
        while left < right:
            mid = int(left + (right - left) / 2)
            if isBadVersion(mid):
                right = mid
            else:
                min_version = mid + 1
                left = mid + 1
        return min_version
