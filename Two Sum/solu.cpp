#include <vector>
#include <iostream>
using namespace std;
class Solution {
public:
	vector<int> twoSum(vector<int>& nums, int target) { //brute-force
		vector<int> result;
		for (int i = 1; i < nums.size(); i++)
			for (int j = i + 1; j < nums.size(); j++)
				if (nums[i] + nums[j] == target)					
					return vector<int>{i, j};				
	}
};
int main()
{
	vector<int> vec = { 1,2,3,4,5 };
	Solution sol;
	vector<int> res = sol.twoSum(vec, 8);
	cout << res[0] << res[1];
	return 0;
}