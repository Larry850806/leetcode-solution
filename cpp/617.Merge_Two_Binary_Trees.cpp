// https://leetcode.com/problems/merge-two-binary-trees/description/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
 public:
  // t1 and t2 can be NULL
  TreeNode *mergeTrees(TreeNode *t1, TreeNode *t2) {
    if (t1 == NULL && t2 == NULL) {
      return NULL;
    }
    if (t1 == NULL) {
      return t2;
    }
    if (t2 == NULL) {
      return t1;
    }

    TreeNode *leftTree = mergeTrees(t1->left, t2->left);
    TreeNode *rightTree = mergeTrees(t1->right, t2->right);
    int rootValue = t1->val + t2->val;
    TreeNode *mergedTree = new TreeNode(rootValue);
    mergedTree->left = leftTree;
    mergedTree->right = rightTree;
    return mergedTree;
  }
};

// class Solution {
//  public:
//   TreeNode *mergeTrees(TreeNode *t1, TreeNode *t2) {
//     if (!t2) return t1;
//     if (!t1) return t2;
//     t1->val += t2->val;
//     t1->left = mergeTrees(t1->left, t2->left);
//     t1->right = mergeTrees(t1->right, t2->right);
//     return t1;
//   }
// };