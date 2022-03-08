class Solution {
    
    func find(_ target: Int, from: ArraySlice<Int>) -> Bool {
        if from.count <= 0 { return true }
        if target < 0 { return false }
        
        var v = 0, range = from.indices
        for it in from.enumerated() {
            v = v*10 + it.element
            range = range.dropFirst()
            
            if v > target { return false }
            if v < target { continue }
            
            if find(target-1, from: from[range]) {
                return true
            }
        }
        
        return false
    }
    
    func splitString(_ s: String) -> Bool {
        let nums = s.map { c in
            Int(c.asciiValue ?? 48) - 48
        }
        
        var v = 0, len = 0, range = nums.indices
        for it in nums.enumerated() {
            if it.offset + 1 >= nums.count { break }
            v = v*10 + it.element
            
            len += (v != 0) ? 1 : 0
            if nums.count - it.offset < len-1 { return false }
            
            range = range.dropFirst()
            if find(v-1, from: nums[range]) {
                return true
            }
        }
        
        return false
    }
}