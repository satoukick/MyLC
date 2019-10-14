#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
public:
	vector<string> letterCombinations(string digits) {
		vector<string> res;
		if (digits.empty()) return vector<string>();
		static const vector<string> v = { "","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz" };
		res.push_back("");
		for (int i = 0; i < digits.size(); i++) {
			int num = digits[i] - '0';
			if (num < 0 || num>9) break;
			vector<string> tmp;
			const string candicate = v[num];
			if (candicate.length() == 0) continue;
			for(int j = 0; j< candicate.size();j++)
				for (int k = 0; k < res.size(); k++)
					tmp.push_back(res[k] + candicate[j]);
			res.swap(tmp);

		}
		return res;

	}
};

int main()
{
	Solution solu;
	vector<string> res = solu.letterCombinations("57");
	/*
	for (vector<string>::iterator iter = res.begin(); iter != res.end(); ++iter)
		cout << *iter << endl;
	*/
	for (int i = 0; i < res.size(); ++i)
		cout << res[i] << endl;


}
