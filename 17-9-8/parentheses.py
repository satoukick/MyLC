class Solution(object):
    def generateParenthesis(self, n):
        def generator(p,left,right):
            if right>=left>=0:
                if not right:
                    yield p
                for q in generator(p+'(',left+1,right): yield q
                for q in generator(p+')',left,right-1): yield q
        return list(generator('',n,n))