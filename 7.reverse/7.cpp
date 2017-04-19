#include <iostream>
#include <cmath>
#include <vector>
using namespace std;


class Solution {
public:
	int reverse(int x) {
		if (x == 0) return 0;
		int flag = 0;
		long long int result = 0;;
		if (x < 0)
		{
			flag = 1; x = abs(x);
		}

		vector<int> rev;
		while (x != 0) {
			rev.push_back(x % 10);
			x = x / 10;
		}
		for (auto i = rev.begin(); i != rev.end(); i++)
			result = result * 10 + *i;
		if (flag == 1)
			result *= -1;

		return (result > INT_MAX || result < INT_MIN) ? 0 : result; // handle the overflow of result instead of input
	}
};

int main()
{
	Solution solu;
	cout << solu.reverse(0) << endl;
}