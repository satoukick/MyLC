#include <queue>
#include <vector>
#include <iostream>

using namespace std;


//Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
	vector<vector<int>> levelOrder(TreeNode* root) {
		vector<vector<int>> result;
		if (!root)
			return result;

		queue<TreeNode*> empty, currentQueue, nextQueue;
		result.push_back(vector<int>(1,root->val));

		TreeNode* node;
		if(root->left != NULL)
			nextQueue.push(root->left);
		if(root->right != NULL)
			nextQueue.push(root->right);

		while (!nextQueue.empty())
		{
			vector<int> level;
			currentQueue = nextQueue;
			nextQueue = empty;
			while (!currentQueue.empty())
			{
				node = currentQueue.front();
				currentQueue.pop();
				level.push_back(node->val);
				if(node->left != NULL)
					nextQueue.push(node->left);
				if(node->right != NULL)
					nextQueue.push(node->right);			
			}
			result.push_back(level);
		}
		return result;
	}
};