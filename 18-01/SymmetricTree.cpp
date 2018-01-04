#include <queue>
#include <cstdlib>

//Definition for a binary tree node.
 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
   TreeNode(int x) : val(x), left(NULL), right(NULL) {}


class Solution {
public:
	bool isSymmetric(TreeNode* root) {
		if (!root)
			return true;

		std::queue<TreeNode*> q1, q2;
		q1.push(root->left);
		q2.push(root->right);
		TreeNode *leftNode, *rightNode;

		while (!q1.empty() && !q2.empty())
		{
			leftNode = q1.front();
			q1.pop();
			rightNode = q2.front();
			q2.pop();

			if (leftNode == NULL && rightNode == NULL)
				continue;
			if (leftNode == NULL || rightNode == NULL)
				return false;
			if (leftNode->val != rightNode->val)
				return false;

			q1.push(leftNode->left);
			q1.push(leftNode->right);
			q2.push(rightNode->right);
			q2.push(rightNode->left);
		}
		return true;
		
	   }
};