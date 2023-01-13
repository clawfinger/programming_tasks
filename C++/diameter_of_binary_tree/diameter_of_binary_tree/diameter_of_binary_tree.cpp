// https://leetcode.com/problems/diameter-of-binary-tree/description/

#include <iostream>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}    
};

int depth(TreeNode* root) {
    if (root->left == nullptr && root->right == nullptr) {
        return 0;
    }
    int left = 0;
    int right = 0;
    if (root->left != nullptr) {
        left = depth(root->left) + 1;
    }
    if (root->right != nullptr) {
        right = depth(root->right) + 1;
    }
    return std::max(left, right);
}

int treeDiameterImpl(TreeNode* root, int* max) {
    int left = 0;
    int right = 0;
    if (root->left != nullptr) {
        left = depth(root->left) + 1;
    }
    if (root->right != nullptr) {
        right = depth(root->right) + 1;
    }
    int diam = left + right;
    if (diam > *max) {
        *max = diam;
    }
    return left + right;
}

void weNeedToGoDeeper(TreeNode* root, int* max) {
    if (root == nullptr) {
        return;
    }
    treeDiameterImpl(root, max);
    weNeedToGoDeeper(root->left, max);
    weNeedToGoDeeper(root->right, max);
}

int diameterOfBinaryTree(TreeNode* root) {
    int max = 0;
    weNeedToGoDeeper(root, &max);
    return max;
}



int main()
{
    //[2, 3, null, 1]

    TreeNode* root = new TreeNode(2);
    TreeNode* three = new TreeNode(3);
    
    three->left = new TreeNode(1);
    //three->right = new TreeNode(5);
    root->left = three;
    root->right = new TreeNode(4);
    //std::cout << diameterOfBinaryTree(root);
    std::cout << diameterOfBinaryTree(root);
}