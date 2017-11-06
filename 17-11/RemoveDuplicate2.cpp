#include <vector>
using namespace std;

class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		int cur = 0;
		for (auto iter = nums.cbegin(); iter != nums.cend(); iter++) {
			if (cur < 2 || (*iter) > nums[cur - 2])
				nums[cur++] = (*iter);
		}
		return cur;
	}
};