class Solution {
    func getMaxRepetitions(_ s1: String, _ n1: Int, _ s2: String, _ n2: Int) -> Int {
        var c1 = 0, c2 = 0
        let numSeq1 = s1.map { c -> UInt8 in
            let n = (c.asciiValue ?? 0) - 97
            c1 |= 1<<n
            return n
        }, numSeq2 = s2.map { c -> UInt8 in
            let n = (c.asciiValue ?? 0) - 97
            c2 |= 1<<n
            return n
        }
        if c1&c2 != c2 {
            return 0
        }
        
        var next = Array(repeating: -1, count: numSeq1.count)
        for i in 0..<numSeq1.count {
            if i>0 && numSeq1[i-1] != numSeq2[0] {
                next[i] = next[i-1]-1
                continue
            }
            
            var len = 0, j = 0
            while j < numSeq2.count {
                if numSeq1[(i+len)%numSeq1.count] == numSeq2[j] {
                    j += 1
                }
                len += 1
            }
            next[i] = len
        }
        
        var len = numSeq1.count*n1, i = 0, opt = 0
        while next[i] <= len {
            len -= next[i]
            i = (i+next[i])%numSeq1.count
            opt += 1
        }
        
        return opt/n2
    }
}