/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
 static auto _ = [](){
    std::ios::sync_with_stdio(false);
    std::cin.tie(NULL);
    return nullptr;
}();

class Solution {
public:
    vector<int> getAllElements(TreeNode* root1, TreeNode* root2) {
        std::ios::sync_with_stdio(false);
        std::cin.tie(NULL);
        
        std::vector<int> opt;
        
        std::stack<TreeNode*> s1, s2;
        if (root1 != NULL) s1.push(root1);
        if (root2 != NULL) s2.push(root2);
        
        auto v1 = midCheck(s1), v2 = midCheck(s2);
        while (v1 != NULL || v2 != NULL) {
            if (v1 == NULL || (v2 != NULL && v2->val < v1->val)) {
                opt.push_back(v2->val);
                delete v2;
                v2 = midCheck(s2);
                continue;
            }
            
            opt.push_back(v1->val);
            delete v1;
            v1 = midCheck(s1);
        }
        
        return opt;
    }
    
    TreeNode* midCheck(std::stack<TreeNode*> &s) {
        while (!s.empty()) {
            auto top = s.top();
            if (top->left != NULL) {
                s.push(top->left);
                top->left = NULL;
                continue;
            }
            
            s.pop();
            if (top->right != NULL) {
                s.push(top->right);
                top->right = NULL;
            }
            return top;
        }
        
        return NULL;
    }
};