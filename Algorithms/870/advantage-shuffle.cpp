bool cmp(const pair<int, int> a, const pair<int, int> b) {
    return a.second > b.second;
}

class Solution {
public:
    vector<int> advantageCount(vector<int>& A, vector<int>& B) {
        vector<pair<int, int>> sortB;
        for (int i = 0; i < B.size(); ++i) {
            sortB.push_back({i, B[i]});
        }
        
        sort(A.begin(), A.end());
        sort(sortB.begin(), sortB.end(), cmp);
        
        for (auto pariB : sortB) {
            int i = pariB.first;
            int numB = pariB.second;
            auto numA = upper_bound(A.begin(), A.end(), numB);
            if (numA == A.end()) {
                B[i] = A[0];
                A.erase(A.begin());
            } else {
                B[i] = *numA;
                A.erase(numA);
            }
        }
        return B;
    }
};