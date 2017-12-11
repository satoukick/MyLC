#include <vector>

class Solution {
public:
	int numTrees(int n) {
		std::vector<int> f;
		f.push_back(1);
		for (int i = 1; i <= n; i++) {
			int g = 0;
			for (int j = 0; j < i; j++)
				g += f[j] * f[i - j - 1];
			f.push_back(g);
		}
		return f.back();
	}
};