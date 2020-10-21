class Solution {
public:
    vector<int> findSwapValues(vector<int>& array1, vector<int>& array2) {
        int dif = 0;
        for (int v : array1) dif += v;
        for (int v : array2) dif -= v;
        if (dif&1) return {};
        
        std::sort(array1.begin(), array1.end());
        std::sort(array2.begin(), array2.end());
        
        for (int iA = 0, iB = 0; iA < array1.size(); ++iA) {
            int b = array1[iA] - dif/2;
            while (iB < array2.size() && array2[iB] < b) ++iB;
            if (iB >= array2.size()) return {};
            if (array2[iB] != b) continue;
            
            return {array1[iA], array2[iB]};
        }
        
        return {};
    }
};