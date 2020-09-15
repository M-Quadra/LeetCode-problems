class Solution {
public:
    int longestOnes(vector<int>& A, int K) {
        int ed = 0, last = A[0];
        int cnt = 0;
        for (int v: A) {
            if (v != last) {
                A[ed++] = cnt;
                cnt = 0;
                last = v;
            }
            
            v == 1 ? ++cnt : --cnt;
            last = v;
        }
        A[ed++] = cnt;
        
        int l = 0, r = 0;
        if (A[0] < 0) l = r = 1;
        int p = A[l], n = 0;
        
        int opt = 0;
        while (r < ed) {
            if (K < -n) {
                p -= A[l];
                n -= l+1 < ed ? A[l+1] : 0;
                l += 2;
                r = std::max(l, r);
                continue;
            }
            
            int v = (l-1 >= 0 ? -A[l-1] : 0) + (r+1 < ed ? -A[r+1] : 0);
            v = std::min(K + n, v) + p - n;
            opt = std::max(opt, v);
            p += r+2 < ed ? A[r+2] : 0;
            n += r+1 < ed ? A[r+1] : 0;
            r += 2;
        }
        
        return opt;
    }
};