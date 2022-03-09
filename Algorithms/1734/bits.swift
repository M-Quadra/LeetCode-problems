class Solution {
    
    func decode(_ encoded: [Int]) -> [Int] {
        var opt = Array(1...encoded.count+1)
        var a = 0, b = 0, c = 0
        for (i, v) in encoded.enumerated() {
            if i & 1 == 0 {
                a ^= v
            } else {
                b ^= v
            }
            c ^= opt[i]
        }
        c ^= opt.last ?? 0
        
        for x in 1...encoded.count+1 {
            let y = a^b^x
            if y == 0 { continue }
            if a ^ y != c { continue }
            
            opt[0] = x
            for (i, v) in encoded.enumerated() {
                opt[i+1] = v ^ opt[i]
            }
            return opt
        }
        return []
    }
}