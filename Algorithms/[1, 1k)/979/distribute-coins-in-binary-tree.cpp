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
    
    int cnt;
    
    int dfs(TreeNode* note) {
        if (note == NULL) return 0;
        
        int optCnt = 0;
        
        int lft = dfs(note->left);
        int rit = dfs(note->right);
        
        note->val -= 1;
        note->val += lft;
        note->val += rit;
        
        if (lft < 0) optCnt -= lft;
        if (rit < 0) optCnt -= rit;
        
        optCnt += max(0, note->val);
        cnt += optCnt;
        return note->val;
    }
    
    int distributeCoins(TreeNode* root) {
        cnt = 0;
        dfs(root);
        return cnt;
    }
};