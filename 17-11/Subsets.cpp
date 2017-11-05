#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
	vector<vector<int>> subsets(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		vector<vector<int>> results;
		vector<int> path;
		gensub_recursive(nums, 0, path, results);
		return results;
	
	}

	void gensub_recursive(vector<int>& nums, int index, vector<int> &sub, vector<vector<int>> &subs) {
		subs.push_back(sub);
		for (int i = index; i < nums.size(); ++i) {
			sub.push_back(nums[i]);
			gensub_recursive(nums, i + 1, sub, subs);
			sub.pop_back();
		}
	}
};

int main()
{

}