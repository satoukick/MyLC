class Solution:
    def countOfAtoms(self, formula):
        """
        :type formula: str
        :rtype: str
        """
        import re
        from collections import defaultdict

        stack = [defaultdict(int)]
        pattern = '([A-Z][a-z]?|\(|\)|\d+)'
        tokens = list(filter(lambda c: c, re.split(pattern, formula)))
        i = 0

        while i < len(tokens):
            token = tokens[i]
            if token == '(':
                stack.append(defaultdict(int))
            else:
                count = 1
                if i + 1 < len(tokens) and re.match('^\d+$', tokens[i+1]):
                    count, i = int(tokens[i+1]), i + 1
                atmos = stack.pop() if token == ')' else {token: 1}
                for atmo in atmos:
                    stack[-1][atmo] += atmos[atmo] * count
            i += 1

        return ''.join([atmo + (str(count) if count > 1 else '') for atmo, count in sorted(stack[-1].items())])
