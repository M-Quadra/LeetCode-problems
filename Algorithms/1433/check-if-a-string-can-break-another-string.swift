class Solution {
    
    func checkIfCanBreak(_ s1: String, _ s2: String) -> Bool {
        var arr0 = s1.reduce(into: [Int](repeating: 0, count: 26)) { arr, char in
            let i = Int(char.asciiValue ?? 97) - 97
            arr[i] += 1
        }
        var arr1 = s2.reduce(into: [Int](repeating: 0, count: 26)) { arr, char in
            let i = Int(char.asciiValue ?? 97) - 97
            arr[i] += 1
        }
        
        var i = 0, j = 0, m0 = false, m1 = false
        while i < 26 && j < 26 {
            while i < 26 && arr0[i] <= 0 { i += 1 }
            while j < 26 && arr1[j] <= 0 { j += 1 }
            if i >= 26 || j >= 26 { break }
            
            let v = min(arr0[i], arr1[j])
            arr0[i] -= v
            arr1[j] -= v
            if i < j {
                m0 = true
            } else if i > j {
                m1 = true
            }
            if m0 && m1 { return false }
        }
        
        return true
    }
}