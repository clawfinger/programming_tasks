// https://leetcode.com/problems/binary-tree-preorder-traversal/

#include <iostream>
#include <vector>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}    
};

void preorderTraversalImpl(std::vector<int>& output, TreeNode* root) {
    if (root == nullptr) {
        return;
    }
    output.push_back(root->val);
    preorderTraversalImpl(output, root->left);
    preorderTraversalImpl(output, root->right);
}

std::vector<int> preorderTraversal(TreeNode * root) {
    std::vector<int> res;
    preorderTraversalImpl(res, root);
    return res;
}

int main()
{
    TreeNode* root = new TreeNode(5);
    auto res = preorderTraversal(root);
    std::cout << "Hello World!\n";
}