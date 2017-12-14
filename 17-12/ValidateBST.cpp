#include <stdio.h>

//Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode *left;     
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
	bool isValidBST(TreeNode* root)
	{
		TreeNode* p = NULL;
		return isvalidate(root, p);
	}

	bool isvalidate(TreeNode* Node, TreeNode* &prev)
	{
		if (Node == NULL) return true; //root
		if (!isvalidate(Node->left, prev)) return false;// validate left-subtree
		if (prev != NULL && prev->val >= Node->val) return false; // judge value
		prev = Node;
		return isvalidate(Node->right, prev);// validate right-subtree
	}
};