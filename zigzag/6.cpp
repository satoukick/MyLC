#include <iostream>
#include <string>
using namespace std;
class Solution { // 26ms
public:
	string convert(string s, int numRows) {
		if (s.size() == 0 || numRows <= 1) return s;
		string *zigzag = new string[numRows];
		for (int i = 0; i < numRows; i++)
			zigzag[i] = "";
		int max = 2 * numRows - 2;
		for (int i = 0; i < s.size(); i++) {
			if (i % max < numRows)
				zigzag[i % max] += s[i];
			else
				zigzag[max - i % max] += s[i];
		}
		string res = "";
		for (int j = 0; j < numRows; j++)
			res += zigzag[j];
		return res;
		}
};

int main()
{
	string input = "ABCDEFGHIJK";
	Solution solt;
	cout << solt.convert(input, 4) << endl;

}
