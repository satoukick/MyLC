#include <iostream>
#include <stack>
#include <queue>
using namespace std;
//Definition for singly-linked list.
struct ListNode {
	int val;
	ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
		ListNode phead(0), *p = &phead;
		int extra = 0, a = 0;
		stack<int> v1, v2, r;
		while (l1)
		{
			v1.push(l1->val);
			l1 = l1->next;
		}
		while (l2)
		{
			v2.push(l2->val);
			l2 = l2->next;
		}
		while (!v1.empty() ||!v2.empty()|| extra != 0)
		{
			int sum = (!v1.empty() ? v1.top() : 0) + (!v2.empty() ? v2.top() : 0) + extra;
			extra = sum / 10;
			r.push(sum % 10);
			if (!v1.empty()) v1.pop();
			if (!v2.empty()) v2.pop();

 		}
		while (!r.empty())
		{
			p->next = new ListNode(r.top());
			p = p->next;
			r.pop();
		}
		return phead.next;
	}
};
