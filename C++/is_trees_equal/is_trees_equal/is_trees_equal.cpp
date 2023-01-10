// https://leetcode.com/problems/same-tree/

#include <iostream>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}
};

bool isSameTreeImpl(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
        return true;
    }
    else if ((p == nullptr && q != nullptr) || (p != nullptr && q == nullptr)) {
        return false;
    }
    if (p->val != q->val) {
        return false;
    }
    bool left = isSameTreeImpl(p->left, q->left);
    if (left != true) {
        return false;
    }
    bool right = isSameTreeImpl(p->right, q->right);
    if (right != true) {
        return false;
    }
    return true;
}

bool isSameTree(TreeNode* p, TreeNode* q) {
    return isSameTreeImpl(p, q);
}

int main()
{
    TreeNode* p = new TreeNode(1);
    p->left = new TreeNode(2);
    p->right = new TreeNode(3);
    TreeNode* q = new TreeNode(1);
    q->left = new TreeNode(2);
    q->right = new TreeNode(3);
    std::cout << isSameTree(p, q);
}
