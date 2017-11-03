class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        places = [p for p in path.split('/') if path.split('/') if p != '.' or p != '']
        stack  =[]
        for p in places:
            if p == "..":
                stack.pop()
            else:
                stack.append(p)

        return  "/" + "/".join(stack)