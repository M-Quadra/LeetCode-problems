class Solution {
    
    func minWindow(_ s: String, _ t: String) -> String {
        var counter = t.reduce(into: [Int](repeating: 0, count: 123-65)) { arr, char in
            arr[Int(char.asciiValue ?? 65) - 65] -= 1
        }
        var cnt = counter.reduce(0) { $0 + ($1 < 0 ? 1 : 0) }
        
        var start = 0, end = 0
        var l = 0
        let arr = s.map { Int($0.asciiValue ?? 65) - 65 }
        
        for (r, v) in arr.enumerated() {
            counter[v] += 1
            if counter[v] == 0 { cnt -= 1 }
            if cnt > 0 { continue }
            
            while l < r && counter[arr[l]] > 0 {
                counter[arr[l]] -= 1
                l += 1
            }
            
            if end != 0 && r-l+1 > end-start { continue }
            (start, end) = (l, r + 1)
        }
        
        let st = s.index(s.startIndex, offsetBy: start)
        let ed = s.index(s.startIndex, offsetBy: end)
        return String(s[st..<ed])
    }
}